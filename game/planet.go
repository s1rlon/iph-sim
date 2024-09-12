package game

import "math"

type Planet struct {
	Name           string
	Ores           []Ore
	Distribution   []float64
	MiningLevel    int
	ShipSpeedLeve1 int
	ShipCargoLevel int
	UnlockCost     int
	ColonyLevel    int
	AlchemyLevel   int
	Distance       float64
	Locked         bool
	Manager        *Manager
}

func NewPlanet(name string, ores []Ore, distribution []float64, unlockCost int, distance float64) *Planet {
	return &Planet{
		Name:           name,
		Ores:           ores,
		Distribution:   distribution,
		MiningLevel:    1, // Default value
		ShipSpeedLeve1: 1, // Default value
		ShipCargoLevel: 1, // Default value
		ColonyLevel:    0,
		AlchemyLevel:   0,
		Locked:         true, // Default value
		UnlockCost:     unlockCost,
		Distance:       distance,
	}
}

func (g *Game) GetPlanetByName(name string) *Planet {
	for _, planet := range g.Planets {
		if planet.Name == name {
			return planet
		}
	}
	return nil
}

func (g *Game) UpdateColonyLevel(planetName string, colonyLevel int) {
	planet := g.GetPlanetByName(planetName)
	if planet != nil {
		planet.ColonyLevel = colonyLevel
		updatePlanetDB(DB, planet)
	}
}

func (p *Planet) getMiningRate(level int) float64 {
	levelFloat := float64(level)
	return GlobalCalcer.planetCalcer.getMiningRate(p, levelFloat)
}

func (p *Planet) Mine(level int) map[string]float64 {
	minedOres := make(map[string]float64)
	if p.Locked && level == p.MiningLevel {
		return minedOres
	}
	miningRate := p.getMiningRate(level)
	for i, ore := range p.Ores {
		minedAmount := miningRate * p.Distribution[i]
		minedOres[ore.Name] = minedAmount
	}
	return minedOres
}

func (p *Planet) getLevelUpgradeCost(level int) float64 {
	levelFloat := float64(level)
	base_cost := (float64(p.UnlockCost) / 20) * math.Pow(1.3, levelFloat-1)
	cost := base_cost * 1
	return cost
}

func (p *Planet) getUpgradeCost() float64 {
	if p.Locked {
		return float64(p.UnlockCost)
	}
	mining_cost := p.getLevelUpgradeCost(p.MiningLevel)
	cargo_cost := 0.0
	initalShipLevel := p.ShipSpeedLeve1
	initalCargoLevel := p.ShipCargoLevel
	for !p.isCargoSufficent(p.MiningLevel + 1) {
		if p.isCargoSizeBetterUpgradeForVolume() {
			cargo_cost += p.getLevelUpgradeCost(p.ShipCargoLevel)
			p.ShipCargoLevel++
		} else {
			cargo_cost += p.getLevelUpgradeCost(p.ShipSpeedLeve1)
			p.ShipSpeedLeve1++
		}
	}
	cost := mining_cost + cargo_cost
	p.ShipSpeedLeve1 = initalShipLevel
	p.ShipCargoLevel = initalCargoLevel
	return cost
}

func (p *Planet) getMinedOresValue(level int) float64 {
	minedOres := p.Mine(level)
	value := 0.0
	for ore, amount := range minedOres {
		value += amount * GetOreValue(ore)
	}
	return value
}

func (p *Planet) getUpgradeROITime() float64 {
	currentValue := p.getMinedOresValue(p.MiningLevel)
	newValue := p.getMinedOresValue(p.MiningLevel + 1)
	valueIncrease := newValue - currentValue
	upgradeCost := p.getUpgradeCost()
	if valueIncrease > 0 {
		return upgradeCost / valueIncrease
	}
	return math.MaxFloat64
}

func (p *Planet) getShipSpeed(level int) float64 {
	levelfloat := float64(level)
	return GlobalCalcer.planetCalcer.getShipSpeed(p, levelfloat)
}

func (p *Planet) getShipCargo(level int) float64 {
	levelfloat := float64(level)
	return GlobalCalcer.planetCalcer.getShipCargo(p, levelfloat)
}

func (p *Planet) getShippingVolume() float64 {
	return p.getShipSpeed(p.ShipSpeedLeve1) * p.getShipCargo(p.ShipCargoLevel) / p.Distance
}

func (p *Planet) isCargoSufficent(level int) bool {
	rate := p.getMiningRate(level)
	volume := p.getShippingVolume()
	return rate < volume
	//return p.getMiningRate(level) > p.getShippingVolume()
}

func (p *Planet) upgradeMining() {
	p.MiningLevel++
	for !p.isCargoSufficent(p.MiningLevel) {
		if p.isCargoSizeBetterUpgradeForVolume() {
			p.ShipCargoLevel++
		} else {
			p.ShipSpeedLeve1++
		}
	}
	updatePlanetDB(DB, p)
}

func (p *Planet) isCargoSizeBetterUpgradeForVolume() bool {
	currentVolume := p.getShippingVolume()
	cargoCost := p.getLevelUpgradeCost(p.ShipCargoLevel)
	p.ShipCargoLevel++
	cargoVolume := p.getShippingVolume()
	p.ShipCargoLevel--
	speedCost := p.getLevelUpgradeCost(p.ShipSpeedLeve1)
	p.ShipSpeedLeve1++
	speedVolume := p.getShippingVolume()
	p.ShipSpeedLeve1--
	cargoIncrease := cargoVolume - currentVolume
	speedCostIncrease := speedVolume - currentVolume
	cargoValue := cargoIncrease / cargoCost
	speedValue := speedCostIncrease / speedCost
	return cargoValue > speedValue
}

func (p *Planet) resetPlanet() {
	p.MiningLevel = 1
	p.ShipSpeedLeve1 = 1
	p.ShipCargoLevel = 1
	p.ColonyLevel = 0
	p.AlchemyLevel = 0
	p.Locked = true
	p.Manager = nil
	resetPlanetDB(DB, p)
}

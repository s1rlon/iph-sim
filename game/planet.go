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
	Distance       int
	Locked         bool
}

func NewPlanet(name string, ores []Ore, distribution []float64, unlockCost int, distance int) *Planet {
	return &Planet{
		Name:           name,
		Ores:           ores,
		Distribution:   distribution,
		MiningLevel:    1, // Default value
		ShipSpeedLeve1: 1, // Default value
		ShipCargoLevel: 1, // Default value
		UnlockCost:     unlockCost,
		Distance:       distance,
		Locked:         true, // Default value
	}
}

func (p *Planet) getMiningRate(level int) float64 {
	levelFloat := float64(level)
	return 0.25 + (0.1 * (levelFloat - 1)) + (0.017 * (levelFloat - 1) * (levelFloat - 1))
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
	return 1 + 0.2*(levelfloat-1) + (1.0/75)*(levelfloat-1)*(levelfloat-1)
}

func (p *Planet) getShipCargo(level int) float64 {
	levelfloat := float64(level)
	return 5 + 2*(levelfloat-1) + 0.1*(levelfloat-1)*(levelfloat-1)
}

func (p *Planet) getShippingVolume() float64 {
	return p.getShipSpeed(p.ShipSpeedLeve1) * p.getShipCargo(p.ShipCargoLevel) / float64(p.Distance)
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

package game

import "math"

type Planet struct {
	Name           string
	Ores           []*Ore
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
	Rover          bool
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

func (p *Planet) Mine(level int) map[*Ore]float64 {
	minedOres := make(map[*Ore]float64)
	if p.Locked && level == p.MiningLevel {
		return minedOres
	}
	miningRate := p.getMiningRate(level)
	for i, ore := range p.Ores {
		minedAmount := miningRate * p.Distribution[i]
		minedOres[ore] = minedAmount
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
		value += amount * MarketSVC.GetOreValue(ore)
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

func NewPlanet(name string, ores []*Ore, distribution []float64, unlockCost int, distance float64) *Planet {
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

func makeNewPlanets(ores []*Ore) []*Planet {

	return []*Planet{
		NewPlanet("Balor", getOres(ores, "Copper"), []float64{1.0}, 100, 15),
		NewPlanet("Drasta", getOres(ores, "Copper", "Iron"), []float64{0.8, 0.2}, 200, 15.4),
		NewPlanet("Anadius", getOres(ores, "Copper", "Iron"), []float64{0.5, 0.5}, 500, 15.8),
		NewPlanet("Dholen", getOres(ores, "Iron", "Lead"), []float64{0.8, 0.2}, 1250, 15.8),
		NewPlanet("Verr", getOres(ores, "Lead", "Iron", "Copper"), []float64{0.5, 0.3, 0.2}, 5000, 16.4),
		NewPlanet("Newton", getOres(ores, "Lead"), []float64{1.0}, 9000, 17.2),
		NewPlanet("Widow", getOres(ores, "Iron", "Copper", "Silica"), []float64{0.4, 0.4, 0.2}, 15000, 19),
		NewPlanet("Acheron", getOres(ores, "Silica", "Copper"), []float64{0.6, 0.4}, 25000, 19.5),
		NewPlanet("Yangtze", getOres(ores, "Silica", "Aluminium"), []float64{0.8, 0.2}, 40000, 19.5),
		NewPlanet("Solveig", getOres(ores, "Aluminium", "Silica", "Lead"), []float64{0.5, 0.3, 0.2}, 75000, 21),
		NewPlanet("Imir", getOres(ores, "Aluminium"), []float64{1.0}, 150000, 22.5),
		NewPlanet("Relic", getOres(ores, "Lead", "Silica", "Silver"), []float64{0.45, 0.35, 0.2}, 250000, 24.5),
		NewPlanet("Nith", getOres(ores, "Silver", "Aluminium"), []float64{0.8, 0.2}, 400000, 26),
		NewPlanet("Batalla", getOres(ores, "Copper", "Iron", "Gold"), []float64{0.4, 0.4, 0.2}, 800000, 29),
		NewPlanet("Micah", getOres(ores, "Gold", "Silver"), []float64{0.5, 0.5}, 1500000, 30.5),
		NewPlanet("Pranas", getOres(ores, "Gold"), []float64{1.0}, 3000000, 32.5),
		NewPlanet("Castellus", getOres(ores, "Aluminium", "Silica", "Diamond"), []float64{0.4, 0.35, 0.25}, 6000000, 33.5),
		NewPlanet("Gorgon", getOres(ores, "Diamond", "Lead"), []float64{0.8, 0.2}, 12000000, 35),
		NewPlanet("Parnitha", getOres(ores, "Gold", "Platinum"), []float64{0.7, 0.3}, 25000000, 38),
		NewPlanet("Orisoni", getOres(ores, "Platinum", "Diamond"), []float64{0.7, 0.3}, 50000000, 40),
		NewPlanet("Theseus", getOres(ores, "Platinum"), []float64{1.0}, 100000000, 44),
		NewPlanet("Zelene", getOres(ores, "Silver", "Titanium"), []float64{0.7, 0.3}, 200000000, 47.5),
		NewPlanet("Han", getOres(ores, "Titanium", "Diamond", "Gold"), []float64{0.7, 0.2, 0.1}, 400000000, 50),
		NewPlanet("Strennus", getOres(ores, "Titanium", "Platinum"), []float64{0.7, 0.3}, 800000000, 55),
		NewPlanet("Osun", getOres(ores, "Aluminium", "Iridium"), []float64{0.6, 0.4}, 1600000000, 58),
		NewPlanet("Ploitari", getOres(ores, "Iridium", "Diamond"), []float64{0.5, 0.5}, 3200000000, 60),
		NewPlanet("Elysta", getOres(ores, "Iridium"), []float64{1.0}, 6400000000, 63),
		NewPlanet("Tikkuun", getOres(ores, "Iridium", "Titanium", "Palladium"), []float64{0.4, 0.35, 0.25}, 12500000000, 67),
		NewPlanet("Satent", getOres(ores, "Palladium", "Titanium"), []float64{0.6, 0.4}, 25000000000, 70),
		NewPlanet("Urla Rast", getOres(ores, "Palladium", "Diamond"), []float64{0.9, 0.1}, 50000000000, 73),
		NewPlanet("Vular", getOres(ores, "Palladium", "Osmium"), []float64{0.7, 0.3}, 100000000000, 75),
		NewPlanet("Nibiru", getOres(ores, "Osmium", "Iridium"), []float64{0.6, 0.4}, 250000000000, 76),
		NewPlanet("Xena", getOres(ores, "Osmium"), []float64{1.0}, 600000000000, 78),
		NewPlanet("Rupert", getOres(ores, "Palladium", "Osmium", "Rhodium"), []float64{0.55, 0.3, 0.15}, 1500000000000, 78),
		NewPlanet("Pax", getOres(ores, "Rhodium", "Platinum"), []float64{0.5, 0.5}, 4000000000000, 80),
		NewPlanet("Ivyra", getOres(ores, "Rhodium"), []float64{1.0}, 10000000000000, 81),
		NewPlanet("Utrits", getOres(ores, "Rhodium", "Inerton"), []float64{0.8, 0.2}, 25000000000000, 82),
		NewPlanet("Doosie", getOres(ores, "Inerton", "Osmium"), []float64{0.5, 0.5}, 62000000000000, 84),
		NewPlanet("Zulu", getOres(ores, "Inerton"), []float64{1.0}, 160000000000000, 84),
		NewPlanet("Unicae", getOres(ores, "Inerton", "Quadium"), []float64{0.8, 0.2}, 400000000000000, 85),
		NewPlanet("Dune", getOres(ores, "Osmium"), []float64{1.0}, 1000000000000000, 87),
	}
}

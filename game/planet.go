package game

import "math"

type Planet struct {
	Name           string   `gorm:"primaryKey"`
	Ores           []*Ore   `gorm:"-"` // Modified by alchemy
	BaseOreNames   []string `gorm:"-"` // Baseline for alchemy resets
	Distribution   []float64 `gorm:"-"` // Ignored by GORM
	MiningLevel    int      `gorm:"default:1"`
	ShipSpeedLeve1 int      `gorm:"default:1"`
	ShipCargoLevel int      `gorm:"default:1"`
	UnlockCost     int      `gorm:"-"` // Ignored by GORM
	ColonyLevel    int      `gorm:"default:0"`
	AlchemyLevel   int      `gorm:"default:0"`
	AlchemizedOreIndex int  `gorm:"default:-1"` // Index of ore that has alchemy applied
	Distance       float64  `gorm:"-"` // Ignored by GORM
	Locked         bool     `gorm:"default:true"`
	Manager        *Manager `gorm:"foreignKey:PlanetName;references:Name"` // HasOne/Optional relationship
	Rover          bool     `gorm:"default:false"`
	game           *Game    `gorm:"-"`
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
		g.SavePlanet(planet)
	}
}

func (g *Game) UpdateAlchemyLevel(planetName string, alchemyLevel int, oreIndex int) {
	planet := g.GetPlanetByName(planetName)
	if planet != nil {
		planet.AlchemyLevel = alchemyLevel
		planet.AlchemizedOreIndex = oreIndex
		g.ApplyAlchemy(planet)
		g.SavePlanet(planet)
	}
}

func (g *Game) ApplyAlchemy(planet *Planet) {
	// Reset to baseline first
	planet.Ores = getOres(g.Ores, planet.BaseOreNames...)

	if planet.AlchemyLevel > 0 && planet.AlchemizedOreIndex >= 0 && planet.AlchemizedOreIndex < len(planet.Ores) {
		currentOre := planet.Ores[planet.AlchemizedOreIndex]
		var nextOre *Ore
		for i, ore := range g.Ores {
			if ore.Name == currentOre.Name && i+planet.AlchemyLevel < len(g.Ores) {
				nextOre = g.Ores[i+planet.AlchemyLevel]
				break
			}
		}
		if nextOre != nil {
			planet.Ores[planet.AlchemizedOreIndex] = nextOre
		}
	}
}

func (p *Planet) getMiningRate(level int) float64 {
	levelFloat := float64(level)
	return p.game.Calcer.planetCalcer.getMiningRate(p, levelFloat)
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
	cost := base_cost - (base_cost * (1 - p.game.Calcer.getGobalUpgradeCostRedux()))
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
		value += amount * ore.getValue()
	}
	return value
}

func (p *Planet) getUpgradeROITime() float64 {
	currentValue := p.getMinedOresValue(p.MiningLevel)
	newValue := p.getMinedOresValue(p.MiningLevel + 1)
	upgradeCost := p.getUpgradeCost()
	if p.Locked {
		currentValue = p.getMinedOresValue(p.MiningLevel + 9)
		newValue = p.getMinedOresValue(p.MiningLevel + 10)
	}
	valueIncrease := newValue - currentValue
	if valueIncrease > 0 {
		return upgradeCost / valueIncrease
	}
	return math.MaxFloat64
}

func (p *Planet) getShipSpeed(level int) float64 {
	levelfloat := float64(level)
	return p.game.Calcer.planetCalcer.getShipSpeed(p, levelfloat)
}

func (p *Planet) getShipCargo(level int) float64 {
	levelfloat := float64(level)
	return p.game.Calcer.planetCalcer.getShipCargo(p, levelfloat)
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

func (p *Planet) upgradeMining(g *Game) {
	p.MiningLevel++
	for !p.isCargoSufficent(p.MiningLevel) {
		if p.isCargoSizeBetterUpgradeForVolume() {
			p.ShipCargoLevel++
		} else {
			p.ShipSpeedLeve1++
		}
	}
	g.SavePlanet(p)
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

func (p *Planet) resetPlanet(g *Game) {
	p.MiningLevel = 1
	p.ShipSpeedLeve1 = 1
	p.ShipCargoLevel = 1
	p.ColonyLevel = 0
	p.AlchemyLevel = 0
	p.AlchemizedOreIndex = -1
	p.Locked = true
	p.Manager = nil
	p.Rover = false
	g.SavePlanet(p)
}

func NewPlanet(g *Game, name string, ores []*Ore, oreNames []string, distribution []float64, unlockCost int, distance float64) *Planet {
	return &Planet{
		game:           g,
		Name:           name,
		Ores:           ores,
		BaseOreNames:   oreNames,
		Distribution:   distribution,
		MiningLevel:    1, // Default value
		ShipSpeedLeve1: 1, // Default value
		ShipCargoLevel: 1, // Default value
		ColonyLevel:    0,
		AlchemyLevel:   0,
		AlchemizedOreIndex: -1,
		Locked:         true, // Default value
		UnlockCost:     unlockCost,
		Distance:       distance,
	}
}



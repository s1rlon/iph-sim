package game

import "math"

type Planet struct {
	Name         string
	Ores         []Ore
	Distribution []float64
	MiningLevel  int
	UnlockCost   int
	Distance     int
	Locked       bool
}

func NewPlanet(name string, ores []Ore, distribution []float64, unlockCost int, distance int) *Planet {
	return &Planet{
		Name:         name,
		Ores:         ores,
		Distribution: distribution,
		MiningLevel:  1, // Default value
		UnlockCost:   unlockCost,
		Distance:     distance,
		Locked:       true, // Default value
	}
}

func (p *Planet) getMiningRate() float64 {
	level := float64(p.MiningLevel)
	return 0.25 + (0.1 * (level - 1)) + (0.017 * (level - 1) * (level - 1))
}

func (p *Planet) Mine() map[string]float64 {
	minedOres := make(map[string]float64)
	if p.Locked {
		return minedOres
	}
	miningRate := p.getMiningRate()
	for i, ore := range p.Ores {
		minedAmount := miningRate * p.Distribution[i]
		minedOres[ore.Name] = minedAmount
	}
	return minedOres
}

func (p *Planet) getUpgradeCost() float64 {
	if p.Locked {
		return float64(p.UnlockCost)
	}
	level := float64(p.MiningLevel)
	return (float64(p.UnlockCost) / 20) * math.Pow(1.3, level-1)
}

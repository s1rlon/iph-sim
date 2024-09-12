package game

import "math"

func (g *Game) bestUpgradeValue() (*Planet, float64, float64) {
	var bestPlanet *Planet
	bestROI := math.MaxFloat64
	var bestValueIncrease float64

	maxRange := g.Projects.telescopeRange()

	for i, planet := range g.Planets {
		if i >= maxRange {
			break
		}
		ROItime := planet.getUpgradeROITime()
		if ROItime < bestROI {
			bestROI = ROItime
			bestPlanet = planet
			bestValueIncrease = planet.getMinedOresValue(planet.MiningLevel+1) - planet.getMinedOresValue(planet.MiningLevel)
		}
	}

	return bestPlanet, bestROI, bestValueIncrease
}

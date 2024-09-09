package game

import "math"

func BestUpgradeValue(game *Game) (*Planet, float64, float64) {
	var bestPlanet *Planet
	bestROI := math.MaxFloat64
	var bestValueIncrease float64

	for _, planet := range game.Planets {

		ROItime := planet.getUpgradeROITime()
		if ROItime < bestROI {
			bestROI = ROItime
			bestPlanet = planet
			bestValueIncrease = planet.getMinedOresValue(planet.MiningLevel+1) - planet.getMinedOresValue(planet.MiningLevel)
		}
	}

	return bestPlanet, bestROI, bestValueIncrease
}

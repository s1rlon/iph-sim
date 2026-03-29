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
		
		// Calculate value increase consistently with getUpgradeROITime
		currentLevel := planet.MiningLevel
		if planet.Locked {
			currentLevel += 9
		}
		valueIncrease := planet.getMinedOresValue(currentLevel+1) - planet.getMinedOresValue(currentLevel)

		if ROItime < bestROI {
			bestROI = ROItime
			bestPlanet = planet
			bestValueIncrease = valueIncrease
		}
		
		// If we encounter a locked planet, we stop looking at further planets
		// as we can only unlock them in order.
		if planet.Locked {
			break
		}
	}

	return bestPlanet, bestROI, bestValueIncrease
}

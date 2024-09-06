package game

func BestUpgradeValue(game *Game) (*Planet, float64, float64) {
	var bestPlanet *Planet
	bestValue := 0.0
	var bestValueIncrease float64

	for _, planet := range game.Planets {
		var currentValue, newValue, valueIncrease float64

		if planet.Locked {
			currentValue = 0.0
			planet.Locked = false
			newMinedOres := planet.Mine()
			planet.Locked = true
			newValue = 0.0
			for ore, amount := range newMinedOres {
				newValue += amount * OreValues[ore]
			}
			valueIncrease = newValue - currentValue
		} else {
			// Calculate current value of ores mined
			currentMinedOres := planet.Mine()
			currentValue = 0.0
			for ore, amount := range currentMinedOres {
				currentValue += amount * OreValues[ore]
			}

			// Calculate new value of ores mined after upgrading
			planet.MiningLevel++
			newMinedOres := planet.Mine()
			newValue = 0.0
			for ore, amount := range newMinedOres {
				newValue += amount * OreValues[ore]
			}
			planet.MiningLevel--

			valueIncrease = newValue - currentValue

		}
		upgradeCost := planet.getUpgradeCost()
		if upgradeCost > 0 {
			valueToCostRatio := valueIncrease / upgradeCost
			if valueToCostRatio > bestValue {
				bestValue = valueToCostRatio
				bestPlanet = planet
				bestValueIncrease = valueIncrease
			}
		}
	}

	return bestPlanet, bestValue, bestValueIncrease
}

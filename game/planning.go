package game

func BestUpgradeValue(game *Game) (*Planet, float64) {
	var bestPlanet *Planet
	bestValue := 0.0

	for _, planet := range game.Planets {
		// Calculate current value of ores mined
		currentMinedOres := planet.Mine()
		currentValue := 0.0
		for ore, amount := range currentMinedOres {
			currentValue += amount * OreValues[ore]
		}

		// Calculate new value of ores mined after upgrading
		planet.MiningLevel++
		newMinedOres := planet.Mine()
		newValue := 0.0
		for ore, amount := range newMinedOres {
			newValue += amount * OreValues[ore]
		}
		planet.MiningLevel--

		valueIncrease := newValue - currentValue
		upgradeCost := planet.getUpgradeCost()

		if upgradeCost > 0 {
			valueToCostRatio := valueIncrease / upgradeCost
			if valueToCostRatio > bestValue {
				bestValue = valueToCostRatio
				bestPlanet = &planet
			}
		}
	}

	return bestPlanet, bestValue
}

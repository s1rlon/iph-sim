package game

func SimulateUpgrades(game *Game, steps int) {
	// Update the last steps
	game.LastSteps = steps

	// Simulate the specified number of best value upgrades
	for i := 0; i < steps; i++ {
		bestPlanet, bestROI, valueIncrease := game.bestUpgradeValue()
		if bestPlanet != nil {
			if bestPlanet.Locked {
				game.GamdeData.TotalMoneySpent += bestPlanet.getUpgradeCost()
				game.GamdeData.UpgradeHistory = append(game.GamdeData.UpgradeHistory, UpgradeHistory{game.GamdeData.CurrentStep, bestPlanet.Name, bestPlanet.getUpgradeCost(), bestROI, valueIncrease, game.GamdeData.TotalMoneySpent})
				bestPlanet.Locked = false
			} else {
				game.GamdeData.TotalMoneySpent += bestPlanet.getUpgradeCost()
				game.GamdeData.UpgradeHistory = append(game.GamdeData.UpgradeHistory, UpgradeHistory{game.GamdeData.CurrentStep, bestPlanet.Name, bestPlanet.getUpgradeCost(), bestROI, valueIncrease, game.GamdeData.TotalMoneySpent})
				bestPlanet.upgradeMining()
			}
		}
		game.GamdeData.CurrentStep++
	}
}

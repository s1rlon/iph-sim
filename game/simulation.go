package game

import "fmt"

func SimulateUpgrades(game *Game, steps int) {
	// Update the last steps
	game.LastSteps = steps

	// Simulate the specified number of best value upgrades
	for i := 0; i < steps; i++ {
		bestPlanet, bestROI, valueIncrease := game.bestUpgradeValue()
		if bestPlanet != nil {
			if bestPlanet.Locked {
				game.GamdeData.UpgradeHistory = append(game.GamdeData.UpgradeHistory, UpgradeHistory{len(game.GamdeData.UpgradeHistory) + 1, bestPlanet.Name, bestPlanet.getUpgradeCost(), bestROI, valueIncrease, game.moneySpent() + bestPlanet.getUpgradeCost()})
				bestPlanet.Locked = false
			} else {
				game.GamdeData.UpgradeHistory = append(game.GamdeData.UpgradeHistory, UpgradeHistory{len(game.GamdeData.UpgradeHistory) + 1, bestPlanet.Name, bestPlanet.getUpgradeCost(), bestROI, valueIncrease, game.moneySpent() + bestPlanet.getUpgradeCost()})
				bestPlanet.upgradeMining()
				game.GamdeData.UpgradeHistory[len(game.GamdeData.UpgradeHistory)-1].Planet = fmt.Sprintf("%s (%d/%d/%d)", bestPlanet.Name, bestPlanet.MiningLevel, bestPlanet.ShipSpeedLeve1, bestPlanet.ShipCargoLevel)
			}
		}
	}
}

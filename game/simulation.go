package game

import "log"

func SimulateUpgrades(game *Game, steps int) {
	// Update the last steps
	game.LastSteps = steps
	game.resetPlanets()

	// Simulate the specified number of best value upgrades
	for i := 0; i < steps; i++ {
		bestPlanet, bestValue, _ := BestUpgradeValue(game)
		if bestPlanet != nil {
			if bestPlanet.Locked {
				log.Printf("Upgrade %d: Unlocking planet: %s with unlock cost: %.2f", i+1, bestPlanet.Name, bestPlanet.getUpgradeCost())
				game.TotalMoneySpent += bestPlanet.getUpgradeCost()
				bestPlanet.Locked = false
			} else {
				log.Printf("Upgrade %d: Best planet to upgrade: %s with value-to-cost ratio: %.2f", i+1, bestPlanet.Name, bestValue)
				game.TotalMoneySpent += bestPlanet.getUpgradeCost()
				bestPlanet.upgradeMining()
			}
		}
	}
}

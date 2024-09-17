package game

import (
	"fmt"
	"log"
)

func SimulateUpgrades(game *Game, steps int) {
	// Update the last steps
	game.LastSteps = steps

	// Simulate the specified number of best value upgrades
	for i := 0; i < steps; i++ {
		bestPlanet, bestROI, valueIncrease := game.bestUpgradeValue()
		if bestPlanet != nil {
			upgradeHistory := UpgradeHistory{
				Stepnum:       len(game.GameData.UpgradeHistory) + 1,
				Planet:        bestPlanet.Name,
				Upgradecost:   bestPlanet.getUpgradeCost(),
				Roitime:       bestROI,
				ValueIncrease: valueIncrease,
				TotalSpend:    game.moneySpent() + bestPlanet.getUpgradeCost(),
			}
			if bestPlanet.Locked {
				bestPlanet.Locked = false
			} else {
				bestPlanet.upgradeMining()
				upgradeHistory.Planet = fmt.Sprintf("%s (%d/%d/%d)", bestPlanet.Name, bestPlanet.MiningLevel, bestPlanet.ShipSpeedLeve1, bestPlanet.ShipCargoLevel)
			}
			game.GameData.UpgradeHistory = append(game.GameData.UpgradeHistory, upgradeHistory)
			err := upgradeHistory.saveToDB(game.db)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

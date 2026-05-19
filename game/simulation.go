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
			// Capture state BEFORE applying the upgrade/unlock for the history
			upgradeHistory := UpgradeHistory{
				Stepnum:       len(game.GameData.UpgradeHistory) + 1,
				PlanetName:    bestPlanet.Name,
				Upgradecost:   bestPlanet.getUpgradeCost(),
				Roitime:       bestROI,
				ValueIncrease: valueIncrease,
				TotalSpend:    game.moneySpent() + bestPlanet.getUpgradeCost(),
				GameDataID:    game.GameData.ID,
			}

			if bestPlanet.Locked {
				bestPlanet.Locked = false
				// After unlocking, we should probably initialize its levels if they aren't already 1
				// but usually they are default 1.
				game.SavePlanet(bestPlanet)
				// PlanetName remains just the name for an unlock entry
			} else {
				bestPlanet.upgradeMining(game)
				upgradeHistory.PlanetName = fmt.Sprintf("%s (%d/%d/%d)", bestPlanet.Name, bestPlanet.MiningLevel, bestPlanet.ShipSpeedLeve1, bestPlanet.ShipCargoLevel)
			}
			game.GameData.UpgradeHistory = append(game.GameData.UpgradeHistory, upgradeHistory)
			err := game.SaveUpgradeHistory(&upgradeHistory)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

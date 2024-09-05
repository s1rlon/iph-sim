package game

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func GenerateHTMLTable(game *Game) string {
	// Simulate 20 best value upgrades
	for i := 0; i < 20; i++ {
		bestPlanet, bestValue := BestUpgradeValue(game)
		if bestPlanet != nil {
			log.Printf("Upgrade %d: Best planet to upgrade: %s with value-to-cost ratio: %.2f", i+1, bestPlanet.Name, bestValue)
			bestPlanet.MiningLevel++
		}
	}

	// Collect mining data after upgrades
	miningData := make(map[string]map[string]float64)
	uniqueOres := make(map[string]bool)
	totalMined := make(map[string]float64)
	totalPerPlanet := make(map[string]float64)
	totalValuePerPlanet := make(map[string]float64)
	totalValue := 0.0

	for _, planet := range game.Planets {
		minedOres := planet.Mine()
		miningData[planet.Name] = minedOres
		for ore, amount := range minedOres {
			uniqueOres[ore] = true
			totalMined[ore] += amount
			totalPerPlanet[planet.Name] += amount
			totalValuePerPlanet[planet.Name] += amount * OreValues[ore]
			totalValue += amount * OreValues[ore]
		}
	}

	var sortedOres []OreValue
	for ore := range uniqueOres {
		sortedOres = append(sortedOres, OreValue{Name: ore, Value: OreValues[ore]})
	}
	sort.Slice(sortedOres, func(i, j int) bool {
		return sortedOres[i].Value < sortedOres[j].Value
	})

	// Build HTML table
	var sb strings.Builder
	sb.WriteString("<html><head><style>")
	sb.WriteString("body { background-color: #121212; color: #e0e0e0; font-family: Arial, sans-serif; }")
	sb.WriteString("table { width: 100%; border-collapse: collapse; }")
	sb.WriteString("th, td { border: 1px solid #444; padding: 8px; text-align: left; }")
	sb.WriteString("th { background-color: #333; }")
	sb.WriteString("tr:nth-child(even) { background-color: #222; }")
	sb.WriteString("</style></head><body><table>")

	// Table header
	sb.WriteString("<tr><th>Ore</th>")
	for _, planet := range game.Planets {
		sb.WriteString(fmt.Sprintf("<th>%s (%d)</th>", planet.Name, planet.MiningLevel))
	}
	sb.WriteString("<th>Total</th></tr>")

	// Total value per planet
	sb.WriteString("<tr><th>Total Value</th>")
	for _, planet := range game.Planets {
		sb.WriteString(fmt.Sprintf("<td>%s</td>", formatNumber(totalValuePerPlanet[planet.Name])))
	}
	sb.WriteString(fmt.Sprintf("<td>%s</td></tr>", formatNumber(totalValue)))

	// Total mined per planet
	sb.WriteString("<tr><th>Total</th>")
	for _, planet := range game.Planets {
		sb.WriteString(fmt.Sprintf("<td>%s</td>", formatNumber(totalPerPlanet[planet.Name])))
	}
	sb.WriteString("<td></td></tr>")

	// Ore rows
	for _, ore := range sortedOres {
		sb.WriteString(fmt.Sprintf("<tr><th>%s</th>", ore.Name))
		for _, planet := range game.Planets {
			if amount, ok := miningData[planet.Name][ore.Name]; ok {
				sb.WriteString(fmt.Sprintf("<td>%s</td>", formatNumber(amount)))
			} else {
				sb.WriteString("<td>-</td>")
			}
		}
		sb.WriteString(fmt.Sprintf("<td>%s</td></tr>", formatNumber(totalMined[ore.Name])))
	}

	// Upgrade cost per planet
	sb.WriteString("<tr><th>Upgrade Cost</th>")
	for _, planet := range game.Planets {
		sb.WriteString(fmt.Sprintf("<td>%s</td>", formatNumber(planet.getUpgradeCost())))
	}
	sb.WriteString("<td></td></tr>")

	sb.WriteString("</table></body></html>")
	return sb.String()
}

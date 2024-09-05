package game

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func GenerateHTMLTable(game *Game) string {

	bestPlanet, bestValue := BestUpgradeValue(game)
	log.Printf("Best planet to upgrade: %s with value-to-cost ratio: %.2f", bestPlanet.Name, bestValue)

	// Collect mining data
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
		sb.WriteString(fmt.Sprintf("<th>%s(%d)</th>", planet.Name, planet.MiningLevel))

	}
	sb.WriteString("<th>Total</th></tr>")

	// Total value per planet
	sb.WriteString("<tr><th>Total Value</th>")
	for _, planet := range game.Planets {
		sb.WriteString(fmt.Sprintf("<td>%.2f</td>", totalValuePerPlanet[planet.Name]))
	}
	sb.WriteString(fmt.Sprintf("<td>%.2f</td></tr>", totalValue))

	// Total mined per planet
	sb.WriteString("<tr><th>Total</th>")
	for _, planet := range game.Planets {
		sb.WriteString(fmt.Sprintf("<td>%.2f</td>", totalPerPlanet[planet.Name]))
	}
	sb.WriteString("<td></td></tr>")

	// Table rows
	for _, ore := range sortedOres {
		sb.WriteString(fmt.Sprintf("<tr><td>%s</td>", ore.Name))
		for _, planet := range game.Planets {
			if amount, exists := miningData[planet.Name][ore.Name]; exists {
				sb.WriteString(fmt.Sprintf("<td>%.2f</td>", amount))
			} else {
				sb.WriteString("<td>-</td>")
			}
		}
		sb.WriteString(fmt.Sprintf("<td>%.2f</td></tr>", totalMined[ore.Name]))
	}

	// Upgrade cost per planet
	sb.WriteString("<tr><th>Upgrade Cost</th>")
	for _, planet := range game.Planets {
		sb.WriteString(fmt.Sprintf("<td>%.2f</td>", planet.getUpgradeCost()))
	}
	sb.WriteString("<td></td></tr>")

	sb.WriteString("</table></body></html>")
	return sb.String()
}

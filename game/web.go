package game

import (
	"fmt"
	"strings"
)

func GenerateHTMLTable(game *Game) string {
	// Define ore values
	oreValues := map[string]float64{
		"Copper":    1.0,
		"Iron":      2.0,
		"Lead":      4.0,
		"Silica":    8.0,
		"Aluminium": 17.0,
		"Silver":    36.0,
		"Gold":      75.0,
		"Diamond":   160.0,
		"Platinium": 340.0,
		"Titanium":  730.0,
		"Iridum":    1600.0,
		"Paladium":  3500.0,
		"Osmium":    7800.0,
		"Rhodium":   17500.0,
		"Inerton":   40000.0,
		"Quadium":   92000.0,
		"Scrith":    215000.0,
		"Uru":       510000.0,
		"Vibranium": 1250000.0,
		"Aether":    3200000.0,
	}

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
			totalValuePerPlanet[planet.Name] += amount * oreValues[ore]
			totalValue += amount * oreValues[ore]
		}
	}

	// Build HTML table
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
		sb.WriteString(fmt.Sprintf("<th>%s</th>", planet.Name))
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
	for ore := range uniqueOres {
		sb.WriteString(fmt.Sprintf("<tr><td>%s</td>", ore))
		for _, planet := range game.Planets {
			if amount, exists := miningData[planet.Name][ore]; exists {
				sb.WriteString(fmt.Sprintf("<td>%.2f</td>", amount))
			} else {
				sb.WriteString("<td>-</td>")
			}
		}
		sb.WriteString(fmt.Sprintf("<td>%.2f</td></tr>", totalMined[ore]))
	}

	sb.WriteString("</table></body></html>")
	return sb.String()
}

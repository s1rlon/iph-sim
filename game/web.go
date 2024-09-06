package game

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"sort"
)

type OreData struct {
	Name    string
	Amounts []float64
	Total   float64
}

type PlanetData struct {
	Name        string
	MiningLevel int
	TotalValue  float64
	TotalMined  float64
	UpgradeCost float64
}

type TableData struct {
	Planets    []PlanetData
	Ores       []OreData
	TotalValue float64
	LastSteps  int
}

func ResetMiningLevels(game *Game) {
	for _, planet := range game.Planets {
		planet.MiningLevel = 1
	}
}

func ResetHandler(game *Game) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResetMiningLevels(game)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func GenerateHTMLTable(game *Game, steps int) string {
	// Update the last steps
	game.LastSteps = steps

	// Simulate the specified number of best value upgrades
	for i := 0; i < steps; i++ {
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

	// Prepare data for template
	var planetData []PlanetData
	for _, planet := range game.Planets {
		planetData = append(planetData, PlanetData{
			Name:        planet.Name,
			MiningLevel: planet.MiningLevel,
			TotalValue:  totalValuePerPlanet[planet.Name],
			TotalMined:  totalPerPlanet[planet.Name],
			UpgradeCost: planet.getUpgradeCost(),
		})
	}

	var oreData []OreData
	for _, ore := range sortedOres {
		var amounts []float64
		for _, planet := range game.Planets {
			if amount, ok := miningData[planet.Name][ore.Name]; ok {
				amounts = append(amounts, amount)
			} else {
				amounts = append(amounts, 0)
			}
		}
		oreData = append(oreData, OreData{
			Name:    ore.Name,
			Amounts: amounts,
			Total:   totalMined[ore.Name],
		})
	}

	data := TableData{
		Planets:    planetData,
		Ores:       oreData,
		TotalValue: totalValue,
		LastSteps:  game.LastSteps,
	}

	// Parse and execute template
	tmpl, err := template.New("planets.html").Funcs(template.FuncMap{
		"formatNumber": formatNumber,
	}).ParseFiles("templates/planets.html")
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	return buf.String()
}

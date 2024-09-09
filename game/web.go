package game

import (
	"bytes"
	"html/template"
	"log"
)

func ResetGalaxy(game *Game) {
	game.resetPlanets()
}

func GenerateHTMLTable(game *Game, steps int) string {
	// Simulate the upgrades
	SimulateUpgrades(game, steps)

	// Create table data
	data := CreateTableData(game)

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

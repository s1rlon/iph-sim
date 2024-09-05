package main

import (
	"fmt"
)

type Ore struct {
	Name string
}

type Planet struct {
	Name         string
	Ores         []Ore
	Distribution []float64
	MiningRate   int
}

type Game struct {
	Planets []Planet
}

func (p *Planet) Mine() map[string]int {
	minedOres := make(map[string]int)
	for i, ore := range p.Ores {
		minedAmount := int(float64(p.MiningRate) * p.Distribution[i])
		minedOres[ore.Name] = minedAmount
	}
	return minedOres
}

func main() {
	// Initialize game with planets
	game := Game{
		Planets: []Planet{
			{
				Name: "Balor",
				Ores: []Ore{
					{Name: "Copper"},
				},
				Distribution: []float64{1.0},
				MiningRate:   20,
			},
			{
				Name: "Drasta",
				Ores: []Ore{
					{Name: "Copper"},
					{Name: "Iron"},
				},
				Distribution: []float64{0.8, 0.2},
				MiningRate:   25,
			},
			{
				Name: "Anadius",
				Ores: []Ore{
					{Name: "Copper"},
					{Name: "Iron"},
				},
				Distribution: []float64{0.5, 0.5},
				MiningRate:   30,
			},
			{
				Name: "Dholen",
				Ores: []Ore{
					{Name: "Iron"},
					{Name: "Lead"},
				},
				Distribution: []float64{0.8, 0.2},
				MiningRate:   35,
			},
		},
	}

	// Collect mining data
	miningData := make(map[string]map[string]int)
	uniqueOres := make(map[string]bool)
	totalMined := make(map[string]int)
	totalPerPlanet := make(map[string]int)

	for _, planet := range game.Planets {
		minedOres := planet.Mine()
		miningData[planet.Name] = minedOres
		for ore, amount := range minedOres {
			uniqueOres[ore] = true
			totalMined[ore] += amount
			totalPerPlanet[planet.Name] += amount
		}
	}

	// Print table header
	fmt.Printf("%-10s", "Ore")
	for _, planet := range game.Planets {
		fmt.Printf("%-10s", planet.Name)
	}
	fmt.Printf("%-10s\n", "Total")

	// Print total mined per planet
	fmt.Printf("%-10s", "Total")
	for _, planet := range game.Planets {
		fmt.Printf("%-10d", totalPerPlanet[planet.Name])
	}
	fmt.Printf("%-10s\n", "")

	// Print table rows
	for ore := range uniqueOres {
		fmt.Printf("%-10s", ore)
		for _, planet := range game.Planets {
			if amount, exists := miningData[planet.Name][ore]; exists {
				fmt.Printf("%-10d", amount)
			} else {
				fmt.Printf("%-10s", "-")
			}
		}
		fmt.Printf("%-10d\n", totalMined[ore])
	}
}

package game

import (
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
	SpeedLevel  int
	CargoLevel  int
	TotalValue  float64
	TotalMined  float64
	UpgradeCost float64
	Locked      bool
}

type TableData struct {
	Planets                  []PlanetData
	Ores                     []OreData
	TotalValue               float64
	LastSteps                int
	NextUpgradePlanet        *PlanetData
	NextUpgradeValueIncrease float64
	TotalMoneySpent          float64
}

func CreateTableData(game *Game) TableData {
	// Determine the next planet to upgrade
	nextPlanet, _, nextValueIncrease := BestUpgradeValue(game)
	var nextUpgradePlanet *PlanetData
	if nextPlanet != nil {
		nextUpgradePlanet = &PlanetData{
			Name:        nextPlanet.Name,
			MiningLevel: nextPlanet.MiningLevel,
			SpeedLevel:  nextPlanet.ShipSpeedLeve1,
			CargoLevel:  nextPlanet.ShipCargoLevel,
			TotalValue:  0, // This can be calculated if needed
			TotalMined:  0, // This can be calculated if needed
			UpgradeCost: nextPlanet.getUpgradeCost(),
			Locked:      nextPlanet.Locked,
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
		minedOres := planet.Mine(planet.MiningLevel)
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
		if !planet.Locked {
			planetData = append(planetData, PlanetData{
				Name:        planet.Name,
				MiningLevel: planet.MiningLevel,
				SpeedLevel:  planet.ShipSpeedLeve1,
				CargoLevel:  planet.ShipCargoLevel,
				TotalValue:  totalValuePerPlanet[planet.Name],
				TotalMined:  totalPerPlanet[planet.Name],
				UpgradeCost: planet.getUpgradeCost(),
				Locked:      planet.Locked,
			})
		}
	}

	var oreData []OreData
	for _, ore := range sortedOres {
		var amounts []float64
		for _, planet := range game.Planets {
			if !planet.Locked {
				if amount, ok := miningData[planet.Name][ore.Name]; ok {
					amounts = append(amounts, amount)
				} else {
					amounts = append(amounts, 0)
				}
			}
		}
		oreData = append(oreData, OreData{
			Name:    ore.Name,
			Amounts: amounts,
			Total:   totalMined[ore.Name],
		})
	}

	return TableData{
		Planets:                  planetData,
		Ores:                     oreData,
		TotalValue:               totalValue,
		LastSteps:                game.LastSteps,
		NextUpgradePlanet:        nextUpgradePlanet,
		NextUpgradeValueIncrease: nextValueIncrease,
		TotalMoneySpent:          game.TotalMoneySpent,
	}
}

package game

import "sort"

type CraftingData struct {
	Name      string
	Type      string
	ProfitInS float64
	Smelters  int
	Crafters  int
	CraftTime float64
	TotalTime float64
}

func (g *Game) MakeCraftingData() []*CraftingData {
	var result []*CraftingData

	availableOres := make(map[string]bool)

	// Determine available ores by mining unlocked planets
	for _, planet := range g.Planets {
		if !planet.Locked {
			minedOres := planet.Mine(planet.MiningLevel)
			for ore := range minedOres {
				availableOres[ore.Name] = true
			}
		}
	}

	// Helper function to check if all inputs are available
	var areAllInputsAvailable func(inputs map[Craftable]float64) bool
	areAllInputsAvailable = func(inputs map[Craftable]float64) bool {
		for input := range inputs {
			if ore, ok := input.(*Ore); ok {
				if !availableOres[ore.Name] {
					return false
				}
			} else if recepie := input.getRecepie(); recepie != nil {
				if !areAllInputsAvailable(recepie.Input) {
					return false
				}
			}
		}
		return true
	}

	for _, recepie := range g.Recepies {
		if !areAllInputsAvailable(recepie.Input) {
			continue
		}
		smelters := recepie.getTotalSmelters()
		crafters := recepie.getTotalCrafters()
		recepieItemTime := recepie.Result.getTime()
		totalTime := recepie.calculateTotalTime(g.GameData.Smelters, g.GameData.Crafters)
		profit := recepie.Result.getValue() / totalTime
		data := &CraftingData{
			Name:      recepie.Result.getName(),
			Type:      recepie.Result.getType(),
			ProfitInS: profit,
			Smelters:  smelters,
			Crafters:  crafters,
			CraftTime: recepieItemTime,
			TotalTime: totalTime,
		}
		result = append(result, data)
	}

	// Sort the result slice by ProfitInS in descending order
	sort.Slice(result, func(i, j int) bool {
		return result[i].ProfitInS > result[j].ProfitInS
	})

	return result
}

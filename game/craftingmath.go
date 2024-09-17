package game

import "sort"

type CraftingData struct {
	Name      string
	Type      string
	ProfitInS float64
	Smelters  int
	Crafters  int
	CraftTime float64
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
	for _, recepie := range g.Recepies {
		allOresAvailable := true
		for input := range recepie.Input {
			if ore, ok := input.(*Ore); ok && !availableOres[ore.Name] {
				allOresAvailable = false
				break
			}
		}
		if !allOresAvailable {
			continue
		}
		profit := recepie.Result.getValue() / recepie.Result.getTime()
		smelters := recepie.getTotalSmelters()
		crafters := recepie.getTotalCrafters()
		recepieItemTime := recepie.Result.getTime()
		data := &CraftingData{
			Name:      recepie.Result.getName(),
			Type:      recepie.Result.getType(),
			ProfitInS: profit,
			Smelters:  smelters,
			Crafters:  crafters,
			CraftTime: recepieItemTime,
		}
		result = append(result, data)
	}

	// Sort the result slice by ProfitInS in descending order
	sort.Slice(result, func(i, j int) bool {
		return result[i].ProfitInS > result[j].ProfitInS
	})

	return result
}

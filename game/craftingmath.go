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
	for _, recepie := range g.Recepies {
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

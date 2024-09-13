package game

type Ore struct {
	Name  string
	Value float64
}

func createOres() *[]*Ore {
	return &[]*Ore{
		{Name: "Copper", Value: 1.0},
		{Name: "Iron", Value: 2.0},
		{Name: "Lead", Value: 4.0},
		{Name: "Silica", Value: 8.0},
		{Name: "Aluminium", Value: 17.0},
		{Name: "Silver", Value: 36.0},
		{Name: "Gold", Value: 75.0},
		{Name: "Diamond", Value: 160.0},
		{Name: "Platinium", Value: 340.0},
		{Name: "Titanium", Value: 730.0},
		{Name: "Iridum", Value: 1600.0},
		{Name: "Paladium", Value: 3500.0},
		{Name: "Osmium", Value: 7800.0},
		{Name: "Rhodium", Value: 17500.0},
		{Name: "Inerton", Value: 40000.0},
		{Name: "Quadium", Value: 92000.0},
		{Name: "Scrith", Value: 215000.0},
		{Name: "Uru", Value: 510000.0},
		{Name: "Vibranium", Value: 1250000.0},
		{Name: "Aether", Value: 3200000.0},
	}
}

func getOres(ores *[]*Ore, names ...string) []*Ore {
	var result []*Ore
	for _, name := range names {
		for _, ore := range *ores {
			if ore.Name == name {
				result = append(result, ore)
				break
			}
		}
	}
	return result
}

func getOre(name string, game *Game) *Ore {
	for _, ore := range *game.Ores {
		if ore.Name == name {
			return ore
		}
	}
	return nil
}

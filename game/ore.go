package game

type Ore struct {
	Name  string
	Value float64
}

type OreValue struct {
	Name  string
	Value float64
}

func createOres() *[]*Ore {
	return &[]*Ore{
		{Name: "Copper", Value: OreValues["Copper"]},
		{Name: "Iron", Value: OreValues["Iron"]},
		{Name: "Lead", Value: OreValues["Lead"]},
		{Name: "Silica", Value: OreValues["Silica"]},
		{Name: "Aluminium", Value: OreValues["Aluminium"]},
		{Name: "Silver", Value: OreValues["Silver"]},
		{Name: "Gold", Value: OreValues["Gold"]},
		{Name: "Diamond", Value: OreValues["Diamond"]},
		{Name: "Platinium", Value: OreValues["Platinium"]},
		{Name: "Titanium", Value: OreValues["Titanium"]},
		{Name: "Iridum", Value: OreValues["Iridum"]},
		{Name: "Paladium", Value: OreValues["Paladium"]},
		{Name: "Osmium", Value: OreValues["Osmium"]},
		{Name: "Rhodium", Value: OreValues["Rhodium"]},
		{Name: "Inerton", Value: OreValues["Inerton"]},
		{Name: "Quadium", Value: OreValues["Quadium"]},
		{Name: "Scrith", Value: OreValues["Scrith"]},
		{Name: "Uru", Value: OreValues["Uru"]},
		{Name: "Vibranium", Value: OreValues["Vibranium"]},
		{Name: "Aether", Value: OreValues["Aether"]},
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

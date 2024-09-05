package game

type Game struct {
	Planets []Planet
}

func NewGame() *Game {
	return &Game{
		Planets: []Planet{
			{
				Name: "Balor",
				Ores: []Ore{
					{Name: "Copper"},
				},
				Distribution: []float64{1.0},
				MiningLevel:  1,
				UnlockCost:   100,
				Distance:     10,
			},
			{
				Name: "Drasta",
				Ores: []Ore{
					{Name: "Copper"},
					{Name: "Iron"},
				},
				Distribution: []float64{0.8, 0.2},
				MiningLevel:  1,
				UnlockCost:   200,
				Distance:     12,
			},
			{
				Name: "Anadius",
				Ores: []Ore{
					{Name: "Copper"},
					{Name: "Iron"},
				},
				Distribution: []float64{0.5, 0.5},
				MiningLevel:  1,
				UnlockCost:   500,
				Distance:     14,
			},
			{
				Name: "Dholen",
				Ores: []Ore{
					{Name: "Iron"},
					{Name: "Lead"},
				},
				Distribution: []float64{0.8, 0.2},
				MiningLevel:  1,
				UnlockCost:   1250,
				Distance:     15,
			},
		},
	}
}

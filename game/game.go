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
				MiningLevel:  2,
			},
			{
				Name: "Drasta",
				Ores: []Ore{
					{Name: "Copper"},
					{Name: "Iron"},
				},
				Distribution: []float64{0.8, 0.2},
				MiningLevel:  1,
			},
			{
				Name: "Anadius",
				Ores: []Ore{
					{Name: "Copper"},
					{Name: "Iron"},
				},
				Distribution: []float64{0.5, 0.5},
				MiningLevel:  4,
			},
			{
				Name: "Dholen",
				Ores: []Ore{
					{Name: "Iron"},
					{Name: "Lead"},
				},
				Distribution: []float64{0.8, 0.2},
				MiningLevel:  1,
			},
		},
	}
}

package game

type Game struct {
	Planets         []*Planet
	LastSteps       int
	TotalMoneySpent float64
}

func NewGame() *Game {
	return &Game{
		Planets: []*Planet{
			NewPlanet("Balor", []Ore{{Name: "Copper"}}, []float64{1.0}, 100, 10),
			NewPlanet("Drasta", []Ore{{Name: "Copper"}, {Name: "Iron"}}, []float64{0.8, 0.2}, 200, 12),
			NewPlanet("Anadius", []Ore{{Name: "Copper"}, {Name: "Iron"}}, []float64{0.5, 0.5}, 500, 14),
			NewPlanet("Dholen", []Ore{{Name: "Iron"}, {Name: "Lead"}}, []float64{0.8, 0.2}, 1250, 15),
		},
		LastSteps: 0,
	}
}

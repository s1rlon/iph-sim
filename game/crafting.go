package game

type Craftable interface {
	GetName() string
}

type Recepie struct {
	Result Craftable
	Input  map[Craftable]float64
}

func createRecepies(game *Game) []*Recepie {

	alloyRecepies := []*Recepie{
		{
			Result: game.getAlloy("Copper Bar"),
			Input: map[Craftable]float64{
				game.getOre("Copper"): 1000,
			},
		},
	}

	itemRecepies := []*Recepie{
		{
			Result: game.getItem("Copper Wire"),
			Input: map[Craftable]float64{
				game.getAlloy("Copper Bar"): 5,
			},
		},
	}
	allRecepies := append(alloyRecepies, itemRecepies...)

	return allRecepies
}

package game

type Craftable interface {
	getName() string
	getBaseValue() float64
	getBaseTime() float64
	getStars() int
	getTrend() float64
	getType() string
	getValue() float64
	getTime() float64
	getRecepie() *Recepie
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
		{
			Result: game.getAlloy("Iron Bar"),
			Input: map[Craftable]float64{
				game.getOre("Iron"): 1000,
			},
		},
		{
			Result: game.getAlloy("Lead Bar"),
			Input: map[Craftable]float64{
				game.getOre("Lead"): 1000,
			},
		},
		{
			Result: game.getAlloy("Silicon Bar"),
			Input: map[Craftable]float64{
				game.getOre("Silica"): 1000,
			},
		},
		{
			Result: game.getAlloy("Aluminium Bar"),
			Input: map[Craftable]float64{
				game.getOre("Aluminium"): 1000,
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
		{
			Result: game.getItem("Iron Nail"),
			Input: map[Craftable]float64{
				game.getAlloy("Iron Bar"): 5,
			},
		},
		{
			Result: game.getItem("Battery"),
			Input: map[Craftable]float64{
				game.getItem("Copper Wire"): 2,
				game.getAlloy("Copper Bar"): 10,
			},
		},
		{
			Result: game.getItem("Hammer"),
			Input: map[Craftable]float64{
				game.getItem("Iron Nail"): 2,
				game.getAlloy("Lead Bar"): 5,
			},
		},
		{
			Result: game.getItem("Glass"),
			Input: map[Craftable]float64{
				game.getAlloy("Silicon Bar"): 10,
			},
		},
	}
	allRecepies := append(alloyRecepies, itemRecepies...)

	return allRecepies
}

func (r *Recepie) getTotalSmelters() int {
	smelters := r.countCraftables("Alloy")
	if r.Result.getType() == "Alloy" {
		smelters++
	}
	return smelters
}

// Method to get the total number of crafters required for a recipe
func (r *Recepie) getTotalCrafters() int {
	crafters := r.countCraftables("Item")
	if r.Result.getType() == "Item" {
		crafters++
	}
	return crafters
}

// Helper method to count the number of craftables of a specific type
func (r *Recepie) countCraftables(craftableType string) int {
	visited := make(map[string]bool)
	return r.countCraftablesRecursive(craftableType, visited)
}

// Recursive method to count the number of craftables of a specific type
func (r *Recepie) countCraftablesRecursive(craftableType string, visited map[string]bool) int {
	count := 0
	for input := range r.Input {
		if !visited[input.getName()] {
			visited[input.getName()] = true
			if input.getType() == craftableType {
				count++
			}
			if input.getRecepie() != nil {
				count += input.getRecepie().countCraftablesRecursive(craftableType, visited)
			}
		}
	}
	return count
}

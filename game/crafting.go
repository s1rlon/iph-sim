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

func (r *Recepie) calculateTotalTime(availableSmelters, availableCrafters int) float64 {
	visited := make(map[string]bool)
	return r.calculateTotalTimeRecursive(availableSmelters, availableCrafters, visited)
}

func (r *Recepie) calculateTotalTimeRecursive(availableSmelters, availableCrafters int, visited map[string]bool) float64 {
	if visited[r.Result.getName()] {
		return 0
	}
	visited[r.Result.getName()] = true

	totalTime := r.Result.getTime()
	for input, quantity := range r.Input {
		inputRecepie := input.getRecepie()
		if inputRecepie != nil {
			inputTime := inputRecepie.calculateTotalTimeRecursive(availableSmelters, availableCrafters, visited)
			if input.getType() == "Alloy" {
				totalTime += inputTime * quantity / float64(availableSmelters)
			} else if input.getType() == "Item" {
				totalTime += inputTime * quantity / float64(availableCrafters)
			}
		}
	}
	return totalTime
}

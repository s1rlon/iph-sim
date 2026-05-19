package game

type Item struct {
	Name     string
	Value    float64
	BaseTime float64
	game     *Game
}

func (i *Item) getName() string {
	return i.Name
}

func (i *Item) getBaseValue() float64 {
	return i.Value
}

func (i *Item) getStars() int {
	return i.game.Market.getStars(i)
}

func (i *Item) getValue() float64 {
	return i.game.Market.getValue(i)
}
func (i *Item) getTrend() float64 {
	return i.game.Market.getTrend(i)
}

func (i *Item) getType() string {
	return "Item"
}
func (i *Item) getBaseTime() float64 {
	return i.BaseTime
}
func (i *Item) getTime() float64 {
	return i.BaseTime
}
func (i *Item) getRecepie() *Recepie {
	return i.game.Market.getRecepieByName(i.Name)
}

func (g *Game) getItem(name string) *Item {
	for _, item := range g.Items {
		if item.Name == name {
			return item
		}
	}
	return nil
}

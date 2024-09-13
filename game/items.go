package game

type Item struct {
	Name     string
	Value    float64
	Recepie  Recepie
	BaseTime float64
}

func (i *Item) getName() string {
	return i.Name
}

func (i *Item) getBaseValue() float64 {
	return i.Value
}

func (i *Item) getStars() int {
	return MarketSVC.getStars(i)
}

func (i *Item) getValue() float64 {
	return MarketSVC.getValue(i)
}

func createItems() []*Item {
	return []*Item{
		{
			Name:     "Copper Wire",
			Value:    10000,
			BaseTime: 60,
		},
	}
}

func (g *Game) getItem(name string) *Item {
	for _, item := range g.Items {
		if item.Name == name {
			return item
		}
	}
	return nil
}

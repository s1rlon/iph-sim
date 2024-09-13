package game

type Item struct {
	Name     string
	Value    float64
	Recepie  Recepie
	BaseTime float64
}

func (i *Item) GetName() string {
	return i.Name
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

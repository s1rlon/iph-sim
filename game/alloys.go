package game

type Alloy struct {
	Name     string
	Value    float64
	Recepie  Recepie
	BaseTime float64
}

func (a *Alloy) getName() string {
	return a.Name
}

func (a *Alloy) getBaseValue() float64 {
	return a.Value
}

func (a *Alloy) getStars() int {
	return MarketSVC.getStars(a)
}

func (a *Alloy) getValue() float64 {
	return MarketSVC.getValue(a)
}

func createAlloys() []*Alloy {
	return []*Alloy{
		{
			Name:     "Copper Bar",
			Value:    1450,
			BaseTime: 20,
		},
	}
}

func (g *Game) getAlloy(name string) *Alloy {
	for _, alloy := range g.Alloys {
		if alloy.Name == name {
			return alloy
		}
	}
	return nil
}

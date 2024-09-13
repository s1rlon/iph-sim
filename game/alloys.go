package game

type Alloy struct {
	Name     string
	Value    float64
	Recepie  Recepie
	BaseTime float64
}

func (a *Alloy) GetName() string {
	return a.Name
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

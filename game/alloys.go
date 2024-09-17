package game

type Alloy struct {
	Name     string
	Value    float64
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

func (a *Alloy) getTrend() float64 {
	return MarketSVC.getTrend(a)
}

func (a *Alloy) getType() string {
	return "Alloy"
}

func (a *Alloy) getBaseTime() float64 {
	return a.BaseTime
}

func (a *Alloy) getTime() float64 {
	return a.BaseTime
}
func (a *Alloy) getRecepie() *Recepie {
	return MarketSVC.getRecepieByName(a.Name)
}

func createAlloys() []*Alloy {
	return []*Alloy{
		{
			Name:     "Copper Bar",
			Value:    1450,
			BaseTime: 20,
		},
		{
			Name:     "Iron Bar",
			Value:    3000,
			BaseTime: 30,
		},
		{
			Name:     "Lead Bar",
			Value:    6010,
			BaseTime: 40,
		},
		{
			Name:     "Silicon Bar",
			Value:    12500,
			BaseTime: 60,
		},
		{
			Name:     "Aluminium Bar",
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

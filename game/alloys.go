package game

type Alloy struct {
	Name     string
	Value    float64
	BaseTime float64
	game     *Game
}

func (a *Alloy) getName() string {
	return a.Name
}

func (a *Alloy) getBaseValue() float64 {
	return a.Value
}

func (a *Alloy) getStars() int {
	return a.game.Market.getStars(a)
}

func (a *Alloy) getValue() float64 {
	return a.game.Market.getValue(a)
}

func (a *Alloy) getTrend() float64 {
	return a.game.Market.getTrend(a)
}

func (a *Alloy) getType() string {
	return "Alloy"
}

func (a *Alloy) getBaseTime() float64 {
	return a.BaseTime
}

func (a *Alloy) getTime() float64 {
	return a.BaseTime / a.game.Calcer.getSmeltSpeedBonus()
}
func (a *Alloy) getRecepie() *Recepie {
	return a.game.Market.getRecepieByName(a.Name)
}

func (g *Game) getAlloy(name string) *Alloy {
	for _, alloy := range g.Alloys {
		if alloy.Name == name {
			return alloy
		}
	}
	return nil
}

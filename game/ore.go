package game

type Ore struct {
	Name  string
	Value float64
	game  *Game
}

func (o *Ore) getName() string {
	return o.Name
}

func (o *Ore) getBaseValue() float64 {
	return o.Value
}

func (o *Ore) getStars() int {
	return o.game.Market.getStars(o)
}

func (o *Ore) getTrend() float64 {
	return o.game.Market.getTrend(o)
}

func (o *Ore) getType() string {
	return "Ore"
}
func (o *Ore) getBaseTime() float64 {
	return 1
}
func (o *Ore) getTime() float64 {
	return 1
}

func (o *Ore) getRecepie() *Recepie {
	return nil
}

func getOres(ores []*Ore, names ...string) []*Ore {
	var result []*Ore
	for _, name := range names {
		for _, ore := range ores {
			if ore.Name == name {
				result = append(result, ore)
				break
			}
		}
	}
	return result
}

func getOre(name string, game *Game) *Ore {
	for _, ore := range game.Ores {
		if ore.Name == name {
			return ore
		}
	}
	return nil
}

func (o *Ore) getValue() float64 {
	return o.game.Market.getValue(o)
}

func (g *Game) getOre(name string) *Ore {
	return getOre(name, g)
}

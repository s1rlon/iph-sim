package game

type Item struct {
	Name     string
	Value    float64
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
func (i *Item) getTrend() float64 {
	return MarketSVC.getTrend(i)
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
	return MarketSVC.getRecepieByName(i.Name)
}

func createItems() []*Item {
	return []*Item{
		{
			Name:     "Copper Wire",
			Value:    10000,
			BaseTime: 60,
		},
		{
			Name:     "Iron Nail",
			Value:    20000,
			BaseTime: 120,
		},
		{
			Name:     "Battery",
			Value:    70000,
			BaseTime: 240,
		},
		{
			Name:     "Hammer",
			Value:    135000,
			BaseTime: 480,
		},
		{
			Name:     "Glass",
			Value:    215000,
			BaseTime: 720,
		},
		{
			Name:     "Circuit",
			Value:    620000,
			BaseTime: 1200,
		},
		{
			Name:     "Lens",
			Value:    1100000,
			BaseTime: 2400,
		},
		{
			Name:     "Laser",
			Value:    3200000,
			BaseTime: 3600,
		},
		{
			Name:     "Basic Computer",
			Value:    7600000,
			BaseTime: 4800,
		},
		{
			Name:     "Solar Panel",
			Value:    15000000,
			BaseTime: 6000,
		},
		{
			Name:     "Laser Torch",
			Value:    31000000,
			BaseTime: 7200,
		},
		{
			Name:     "Advanced Battery",
			Value:    35000000,
			BaseTime: 9000,
		},
		{
			Name:     "Thermal Scanner",
			Value:    71500000,
			BaseTime: 10750,
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

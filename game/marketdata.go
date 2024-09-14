package game

type MarketData struct {
	Ores   []OreData
	Alloys []AlloyData
	Items  []ItemData
}

type OreData struct {
	Name        string
	Value       float64
	Stars       int
	MarketTrend float64
	SellValue   float64
}

type AlloyData struct {
	Name        string
	Value       float64
	Stars       int
	MarketTrend float64
	SellValue   float64
	BaseTime    float64
	CurrentTime float64
}

type ItemData struct {
	Name        string
	Value       float64
	Stars       int
	MarketTrend float64
	SellValue   float64
	BaseTime    float64
	CurrentTime float64
}

func (g *Game) GenerateMarketHTML() MarketData {

	ores := []OreData{}
	for _, ore := range g.Ores {
		ores = append(ores, OreData{
			Name:        ore.Name,
			Value:       ore.Value,
			Stars:       ore.getStars(),
			MarketTrend: ore.getTrend(), // Default market value
			SellValue:   ore.getValue(),
		})
	}

	alloys := []AlloyData{}
	for _, alloy := range g.Alloys {
		alloys = append(alloys, AlloyData{
			Name:        alloy.Name,
			Value:       alloy.Value,
			Stars:       alloy.getStars(), // Default stars
			MarketTrend: alloy.getTrend(), // Default market value
			SellValue:   alloy.getValue(), // Assuming sell value is the same as base value
			BaseTime:    alloy.BaseTime,
			CurrentTime: alloy.BaseTime,
		})
	}

	items := []ItemData{}
	for _, item := range g.Items {
		items = append(items, ItemData{
			Name:        item.Name,
			Value:       item.Value,
			Stars:       item.getStars(), // Default stars
			MarketTrend: item.getTrend(), // Default market value
			SellValue:   item.getValue(), // Assuming sell value is the same as base value
			BaseTime:    item.BaseTime,
			CurrentTime: item.BaseTime,
		})
	}

	data := MarketData{
		Ores:   ores,
		Alloys: alloys,
		Items:  items,
	}
	return data
}

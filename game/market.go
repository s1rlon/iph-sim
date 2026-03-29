package game

type Market struct {
	game  *Game
	Stars map[Craftable]int
	Trend map[Craftable]float64
}

func NewMarket(game *Game) *Market {
	starsMap := make(map[Craftable]int)
	for _, ore := range game.Ores {
		starsMap[ore] = 0
	}
	for _, alloy := range game.Alloys {
		starsMap[alloy] = 0
	}
	for _, item := range game.Items {
		starsMap[item] = 0
	}

	var dbStars []Star
	game.db.Find(&dbStars)
	for _, s := range dbStars {
		item := game.getCratablebyName(s.Name)
		if item != nil {
			starsMap[item] = s.Stars
		}
	}

	trend := make(map[Craftable]float64)
	for _, ore := range game.Ores {
		trend[ore] = 1
	}
	for _, alloy := range game.Alloys {
		trend[alloy] = 1
	}
	for _, item := range game.Items {
		trend[item] = 1
	}

	return &Market{
		game:  game,
		Stars: starsMap,
		Trend: trend,
	}
}

func (m *Market) getValue(item Craftable) float64 {
	value := item.getBaseValue()
	value *= (1 + 0.2*float64(item.getStars()))
	value *= m.getTrend(item)
	return value
}

func (m *Market) getStars(item Craftable) int {
	return m.Stars[item]
}

func (m *Market) getTrend(item Craftable) float64 {
	return m.Trend[item]
}

func (m *Market) saveStars(item Craftable, stars int) error {
	if stars == 0 {
		return m.removeStars(item)
	}
	var star Star
	m.game.db.Where("name = ?", item.getName()).FirstOrCreate(&star, Star{Name: item.getName()})
	star.Stars = stars
	err := m.game.db.Save(&star).Error
	m.Stars[item] = stars
	return err
}

func (m *Market) saveTrend(item Craftable, trend float64) error {
	var err error
	m.Trend[item] = trend
	return err
}

func (m *Market) removeStars(item Craftable) error {
	err := m.game.db.Where("name = ?", item.getName()).Delete(&Star{}).Error
	m.Stars[item] = 0
	return err
}

func (m *Market) getRecepieByName(name string) *Recepie {
	for _, recepie := range m.game.Recepies {
		if recepie.Result.getName() == name {
			return recepie
		}
	}
	return nil
}

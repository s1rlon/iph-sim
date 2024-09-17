package game

type Market struct {
	game  *Game
	Stars map[Craftable]int
	Trend map[Craftable]float64
}

func NewMarket(game *Game) *Market {
	stars := make(map[Craftable]int)
	for _, ore := range game.Ores {
		stars[ore] = 0
	}
	for _, alloy := range game.Alloys {
		stars[alloy] = 0
	}
	for _, item := range game.Items {
		stars[item] = 0
	}
	rows, err := game.db.Query("SELECT name, stars FROM stars")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var name string
		var dbstars int
		if err := rows.Scan(&name, &dbstars); err != nil {
			panic(err)
		}
		item := game.getCratablebyName(name)
		if item != nil {
			stars[item] = dbstars
		}
	}
	defer rows.Close()

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
		Stars: stars,
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
	_, err := m.game.db.Exec("INSERT OR REPLACE INTO stars (name, stars) VALUES (?, ?)", item.getName(), stars)
	m.Stars[item] = stars
	return err
}

func (m *Market) saveTrend(item Craftable, trend float64) error {
	var err error
	m.Trend[item] = trend
	return err
}

func (m *Market) removeStars(item Craftable) error {
	_, err := m.game.db.Exec("DELETE FROM stars WHERE name = ?", item.getName())
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

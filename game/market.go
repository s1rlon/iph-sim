package game

type Market struct {
	game     *Game
	OreStars map[*Ore]int
}

func NewMarket(game *Game) *Market {
	oreStars := make(map[*Ore]int)
	for _, ore := range game.Ores {
		oreStars[ore] = 0
	}
	return &Market{
		game:     game,
		OreStars: oreStars,
	}
}

package game

type Ships struct {
	AdShip       bool
	Daugtership  bool
	Eldership    bool
	Aurora       bool
	Enigma       bool
	Exodus       bool
	Merchant     bool
	Thunderhorse bool
}

func NewShips() *Ships {
	return &Ships{
		AdShip:       false,
		Daugtership:  false,
		Eldership:    false,
		Aurora:       false,
		Enigma:       false,
		Exodus:       false,
		Merchant:     false,
		Thunderhorse: false,
	}
}

func (g *Game) UpdateShips(ships *Ships) {
	g.Ships = ships
}

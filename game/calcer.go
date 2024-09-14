package game

type Calcer struct {
	game         *Game
	planetCalcer *PlanetCalcer
}

type PlanetCalcer struct {
	game *Game
}

func NewCalcer(game *Game) *Calcer {
	return &Calcer{game: game,
		planetCalcer: NewPlanetCalcer(game)}
}

//Global things

//Planet things

func NewPlanetCalcer(game *Game) *PlanetCalcer {
	return &PlanetCalcer{game: game}
}

func (p *PlanetCalcer) getMiningRate(planet *Planet, level float64) float64 {
	rate := 0.25 + (0.1 * (level - 1)) + (0.017 * (level - 1) * (level - 1))
	if planet.Manager != nil {
		if planet.Manager.Primary == Role("Miner") {
			rate *= (1 + 0.25*float64(planet.Manager.Stars))
		}
	}
	//colony level
	rate *= (1 + 0.3*float64(planet.ColonyLevel))
	rate *= p.getMiningGlobalBonus()
	return rate
}

func (p *PlanetCalcer) getMiningGlobalBonus() float64 {
	projects := 1.0
	managers := 1.0
	//Rooms
	rooms := 1 + (0.25 * float64(p.game.Rooms.Engineering))
	//Ships
	ships := 1.0
	if p.game.Ships.AdShip {
		ships += 0.2
	}
	if p.game.Ships.Daugtership {
		ships += 0.5
	}
	if p.game.Ships.Eldership {
		ships += 1.0
	}
	//Station
	station := 1.0
	return projects * managers * rooms * ships * station
}

func (p *PlanetCalcer) getShipSpeed(planet *Planet, level float64) float64 {
	rate := 1 + 0.2*(level-1) + (1.0/75)*(level-1)*(level-1)
	if planet.Manager != nil {
		if planet.Manager.Primary == Role("Pilot") {
			rate *= (1 + 0.25*float64(planet.Manager.Stars))
		}
	}
	return rate
}

func (p *PlanetCalcer) getShipCargo(planet *Planet, level float64) float64 {
	rate := 5 + 2*(level-1) + 0.1*(level-1)*(level-1)
	if planet.Manager != nil {
		if planet.Manager.Primary == Role("Packager") {
			rate *= (1 + 0.5*float64(planet.Manager.Stars))
		}
	}

	return rate
}

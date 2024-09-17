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
			rate *= (1 + 0.25*float64(planet.Manager.Stars)) * p.getGlobalManagerBoost()
		}
	}
	//colony level
	rate *= (1 + 0.3*float64(planet.ColonyLevel))
	//beacon level
	if p.game.Projects.Beacon != 0 {
		rate *= p.getBeaconLevel(planet)
	}
	//global bonus
	rate *= p.getMiningGlobalBonus()
	//projects
	rate *= 1 + (0.25 * float64(p.game.Projects.MiningLevel))
	//station
	rate *= p.game.Station.MineBoost
	return rate
}

func (p *PlanetCalcer) getMiningGlobalBonus() float64 {
	projects := 1.0
	managers := p.getManagerMineBonus()
	//Rooms
	rooms := 1.0
	if p.game.Rooms.Engineering > 0 {
		rooms += 0.25
		if p.game.Rooms.Engineering > 1 {
			rooms += 0.15 * float64(p.game.Rooms.Engineering-1)
		}
	}
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

func (p *PlanetCalcer) getGlobalSpeedBonuus() float64 {
	projects := 1.0
	managers := 1.0
	//Rooms
	rooms := 1.0
	if p.game.Rooms.Aeronautical > 0 {
		rooms += 0.5
		if p.game.Rooms.Aeronautical > 1 {
			rooms += 0.25 * float64(p.game.Rooms.Aeronautical-1)
		}
	}
	//Ships
	ships := 1.0
	if p.game.Ships.Daugtership {
		ships += 0.25
	}
	if p.game.Ships.Eldership {
		ships += 0.5
	}
	//Station
	station := 1.0
	return projects * managers * rooms * ships * station
}

func (p *PlanetCalcer) getShipSpeed(planet *Planet, level float64) float64 {
	rate := 1 + 0.2*(level-1) + (1.0/75)*(level-1)*(level-1)
	if planet.Manager != nil {
		if planet.Manager.Primary == Role("Pilot") {
			rate *= (1 + 0.25*float64(planet.Manager.Stars)) * p.getGlobalManagerBoost()
		}
	}
	rate *= p.getGlobalSpeedBonuus()
	return rate
}

func (p *PlanetCalcer) getGlobalCargoBonuus() float64 {
	projects := 1.0
	managers := 1.0
	//Rooms
	rooms := 1.0
	if p.game.Rooms.Packaging > 0 {
		rooms += 0.5
		if p.game.Rooms.Packaging > 1 {
			rooms += 0.25 * float64(p.game.Rooms.Packaging-1)
		}
	}
	//Ships
	ships := 1.0
	if p.game.Ships.Daugtership {
		ships += 0.25
	}
	if p.game.Ships.Eldership {
		ships += 0.5
	}
	//Station
	station := 1.0
	return projects * managers * rooms * ships * station
}

func (p *PlanetCalcer) getShipCargo(planet *Planet, level float64) float64 {
	rate := 5 + 2*(level-1) + 0.1*(level-1)*(level-1)
	if planet.Manager != nil {
		if planet.Manager.Primary == Role("Packager") {
			rate *= (1 + 0.5*float64(planet.Manager.Stars)) * p.getGlobalManagerBoost()
		}
	}
	rate *= p.getGlobalCargoBonuus()
	return rate
}

func (p *PlanetCalcer) getBeaconLevel(planet *Planet) float64 {
	planetIndex := p.game.getPlanetIndexByName(planet.Name)
	switch {
	case planetIndex <= 3:
		return p.game.Beacon.Levels[0]
	case planetIndex <= 6:
		return p.game.Beacon.Levels[1]
	case planetIndex <= 9:
		return p.game.Beacon.Levels[2]
	case planetIndex <= 12:
		return p.game.Beacon.Levels[3]
	case planetIndex <= 15:
		return p.game.Beacon.Levels[4]
	case planetIndex <= 18:
		return p.game.Beacon.Levels[5]
	case planetIndex <= 21:
		return p.game.Beacon.Levels[6]
	case planetIndex <= 24:
		return p.game.Beacon.Levels[7]
	case planetIndex <= 27:
		return p.game.Beacon.Levels[8]
	case planetIndex <= 30:
		return p.game.Beacon.Levels[9]
	case planetIndex <= 33:
		return p.game.Beacon.Levels[10]
	case planetIndex <= 36:
		return p.game.Beacon.Levels[11]
	case planetIndex <= 39:
		return p.game.Beacon.Levels[12]
	case planetIndex <= 42:
		return p.game.Beacon.Levels[13]
	case planetIndex <= 45:
		return p.game.Beacon.Levels[14]
	case planetIndex <= 48:
		return p.game.Beacon.Levels[15]
	case planetIndex <= 51:
		return p.game.Beacon.Levels[16]
	case planetIndex <= 54:
		return p.game.Beacon.Levels[17]
	case planetIndex <= 57:
		return p.game.Beacon.Levels[18]
	case planetIndex <= 60:
		return p.game.Beacon.Levels[19]
	default:
		return p.game.Beacon.Levels[20]
	}
}

func (c *Calcer) getSmeltSpeedBonus() float64 {
	manager := c.planetCalcer.getManagerSmeltBonus()
	rooms := 1.0
	if c.game.Rooms.Forge > 0 {
		rooms += 0.2
		if c.game.Rooms.Packaging > 1 {
			rooms += 0.1 * float64(c.game.Rooms.Packaging-1)
		}
	}
	return manager * rooms
}

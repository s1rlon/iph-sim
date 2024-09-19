package game

func (p *PlanetCalcer) getGlobalManagerBoost() float64 {
	rooms := 1.0
	if p.game.Rooms.Classroom > 0 {
		rooms += 0.15
		if p.game.Rooms.Classroom > 1 {
			rooms += 0.05 * float64(p.game.Rooms.Classroom-1)
		}
	}
	projects := 1 + (0.1 * float64(p.game.Projects.ManTraining))
	station := p.game.Station.ManagerBoost
	return rooms * projects * station
}

func (p *PlanetCalcer) getGlobalManagerSecondaryBoost() float64 {
	boost := 1 + (0.05 * float64(p.game.Projects.ManSTraing))
	return boost
}

func (p *PlanetCalcer) getManagerMineBonus() float64 {
	bonus := 1.0
	count := 0
	for _, manager := range p.game.GetManagers() {
		if manager.Planet != nil && manager.Secondary == Mine && manager.Stars > 2 {
			bonus += (0.05 * float64(manager.Stars-1))
			count++
		}
	}
	if count > 0 {
		bonus = ((bonus-1)*p.getGlobalManagerBoost()*p.getGlobalManagerSecondaryBoost() + 1)
	}
	return bonus
}

func (p *PlanetCalcer) getManagerSpeedBonus() float64 {
	bonus := 1.0
	count := 0
	for _, manager := range p.game.GetManagers() {
		if manager.Planet != nil && manager.Secondary == Speed && manager.Stars > 2 {
			bonus += (0.05 * float64(manager.Stars-2))
		}
	}
	if count > 0 {
		bonus = ((bonus-1)*p.getGlobalManagerBoost()*p.getGlobalManagerSecondaryBoost() + 1)
	}
	return bonus
}

func (p *PlanetCalcer) getManagerCargoBonus() float64 {
	bonus := 1.0
	count := 0
	for _, manager := range p.game.GetManagers() {
		if manager.Planet != nil && manager.Secondary == Cargo && manager.Stars > 2 {
			bonus += (0.05 * float64(manager.Stars-2))
		}
	}
	if count > 0 {
		bonus = ((bonus-1)*p.getGlobalManagerBoost()*p.getGlobalManagerSecondaryBoost() + 1)
	}
	return bonus
}

func (p *PlanetCalcer) getManagerSmeltBonus() float64 {
	bonus := 1.0
	count := 0
	for _, manager := range p.game.GetManagers() {
		if manager.Planet != nil && manager.Secondary == Smelt && manager.Stars > 2 {
			bonus += (0.05 * float64(manager.Stars-2))
		}
	}
	if count > 0 {
		bonus = ((bonus-1)*p.getGlobalManagerBoost()*p.getGlobalManagerSecondaryBoost() + 1)
	}
	return bonus
}

func (p *PlanetCalcer) getManagerCraftBonus() float64 {
	bonus := 1.0
	count := 0
	for _, manager := range p.game.GetManagers() {
		if manager.Planet != nil && manager.Secondary == Craft && manager.Stars > 2 {
			bonus += (0.05 * float64(manager.Stars-2))
		}
	}
	if count > 0 {
		bonus = ((bonus-1)*p.getGlobalManagerBoost()*p.getGlobalManagerSecondaryBoost() + 1)
	}
	return bonus
}

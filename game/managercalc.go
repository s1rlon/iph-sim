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
	return rooms * projects
}

func (p *PlanetCalcer) getGlobalManagerSecondaryBoost() float64 {
	boost := 1 + (0.05 * float64(p.game.Projects.ManSTraing))
	return boost
}

func (p *PlanetCalcer) getManagerMineBonus() float64 {
	bonus := 1.0
	for _, manager := range p.game.GetManagers() {
		if manager.Planet != nil && manager.Secondary == Mine && manager.Stars > 2 {
			bonus += (0.05 * float64(manager.Stars-2))
		}
	}
	return bonus * p.getGlobalManagerBoost() * p.getGlobalManagerSecondaryBoost()
}

func (p *PlanetCalcer) getManagerSpeedBonus() float64 {
	bonus := 1.0
	for _, manager := range p.game.GetManagers() {
		if manager.Planet != nil && manager.Secondary == Speed && manager.Stars > 2 {
			bonus += (0.05 * float64(manager.Stars-2))
		}
	}
	return bonus * p.getGlobalManagerBoost() * p.getGlobalManagerSecondaryBoost()
}

func (p *PlanetCalcer) getManagerCargoBonus() float64 {
	bonus := 1.0
	for _, manager := range p.game.GetManagers() {
		if manager.Planet != nil && manager.Secondary == Cargo && manager.Stars > 2 {
			bonus += (0.05 * float64(manager.Stars-2))
		}
	}
	return bonus * p.getGlobalManagerBoost() * p.getGlobalManagerSecondaryBoost()
}

func (p *PlanetCalcer) getManagerSmeltBonus() float64 {
	bonus := 1.0
	for _, manager := range p.game.GetManagers() {
		if manager.Planet != nil && manager.Secondary == Smelt && manager.Stars > 2 {
			bonus += (0.05 * float64(manager.Stars-2))
		}
	}
	return bonus * p.getGlobalManagerBoost() * p.getGlobalManagerSecondaryBoost()
}

func (p *PlanetCalcer) getManagerCraftBonus() float64 {
	bonus := 1.0
	for _, manager := range p.game.GetManagers() {
		if manager.Planet != nil && manager.Secondary == Craft && manager.Stars > 2 {
			bonus += (0.05 * float64(manager.Stars-2))
		}
	}
	return bonus * p.getGlobalManagerBoost() * p.getGlobalManagerSecondaryBoost()
}

package game

type Projects struct {
	TelescopeLevel int
	MiningLevel    int
	ShipSpeedLevel int
	ShipCargoLevel int
}

func newProjects() *Projects {
	return &Projects{
		TelescopeLevel: 1,
		MiningLevel:    0,
		ShipSpeedLevel: 0,
		ShipCargoLevel: 0,
	}
}

func (p *Projects) telescopeRange() int {
	return 4 + p.TelescopeLevel*3
}

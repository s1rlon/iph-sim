package game

type Projects struct {
	TelescopeLevel int
	MiningLevel    int
	ShipSpeedLevel int
	ShipCargoLevel int
}

func NewProjects() *Projects {
	return &Projects{
		TelescopeLevel: 0,
		MiningLevel:    0,
		ShipSpeedLevel: 0,
		ShipCargoLevel: 0,
	}
}

func (p *Projects) telescopeRange() int {
	return 4 + p.TelescopeLevel*3
}

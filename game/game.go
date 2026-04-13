package game

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Star struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"uniqueIndex"`
	Stars int
}

type Game struct {
	Planets   []*Planet
	LastSteps int
	Managers  []*Manager
	Projects  *Projects
	db        *gorm.DB
	GameData  *GameData
	Ships     *Ships
	Ores      []*Ore
	Recepies  []*Recepie
	Alloys    []*Alloy
	Items     []*Item
	Rooms     *Rooms
	Beacon    *Beacon
	Station   *Station
}

var GlobalCalcer *Calcer
var DB *gorm.DB
var MarketSVC *Market

func NewGame() *Game {
	db, err := gorm.Open(sqlite.Open("ipm2.sql"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&Planet{},
		&Manager{},
		&Projects{},
		&UpgradeHistory{},
		&Rooms{},
		&Star{},
		&GameData{},
		&Beacon{},
		&Station{},
		&Ships{},
	)
	if err != nil {
		panic(err)
	}

	g := &Game{
		LastSteps: 1,
		db:        db,
	}

	if err := g.loadOres(); err != nil {
		panic(err)
	}
	if err := g.loadPlanets(); err != nil {
		panic(err)
	}
	if err := g.loadAlloys(); err != nil {
		panic(err)
	}
	if err := g.loadItems(); err != nil {
		panic(err)
	}

	g.Managers = g.getManagersFromDB()
	g.Projects = g.loadProjectsFromDB()
	g.GameData = g.loadGameDataFromDB()
	g.Ships = g.loadShipsFromDB()
	g.Rooms = g.loadRoomsFromDB()
	g.Beacon = g.loadBeaconDataFromDB()
	g.Station = g.loadStationDataFromDB()

	return g
}

func (g *Game) InitData() {
	GlobalCalcer = NewCalcer(g)
	MarketSVC = NewMarket(g)
	if err := g.loadRecipes(); err != nil {
		panic(err)
	}
	DB = g.db
	dbPlanets, _ := g.getPlanetsFromDB()
	for _, planet := range g.Planets {
		for _, dbPlanet := range dbPlanets {
			if planet.Name == dbPlanet.Name {
				planet.MiningLevel = dbPlanet.MiningLevel
				planet.ShipSpeedLeve1 = dbPlanet.ShipSpeedLeve1
				planet.ShipCargoLevel = dbPlanet.ShipCargoLevel
				planet.Locked = dbPlanet.Locked
				planet.ColonyLevel = dbPlanet.ColonyLevel
				planet.AlchemyLevel = dbPlanet.AlchemyLevel
				planet.AlchemizedOreIndex = dbPlanet.AlchemizedOreIndex
				g.ApplyAlchemy(planet)
			}
		}
	}
}

func (g *Game) ResetGalaxy() {
	g.ResetPlanets()
	g.ResetManagers()
	g.GameData.resetGameData()
	g.UpdateProjects(newProjects())
}

func (g *Game) ResetPlanets() {
	for _, planet := range g.Planets {
		planet.resetPlanet(g)
	}
}

func (g *Game) ResetManagers() {
	for _, manager := range g.Managers {
		manager.unassignManager(g)
	}
}

func (g *Game) moneySpent() float64 {
	total := 0.0
	for _, upgrade := range g.GameData.UpgradeHistory {
		total += upgrade.Upgradecost
	}
	return total
}

func (g *Game) getCratablebyName(name string) Craftable {
	for _, ore := range g.Ores {
		if ore.getName() == name {
			return ore
		}
	}
	for _, alloy := range g.Alloys {
		if alloy.getName() == name {
			return alloy
		}
	}
	for _, item := range g.Items {
		if item.getName() == name {
			return item
		}
	}
	return nil
}

func (g *Game) SetStars(name string, stars int) {
	item := g.getCratablebyName(name)
	if item != nil {
		MarketSVC.saveStars(item, stars)
	}
}

func (g *Game) SetTrend(name string, trend float64) {
	item := g.getCratablebyName(name)
	if item != nil {
		MarketSVC.saveTrend(item, trend)
	}
}

func (g *Game) getPlanetIndexByName(name string) int {
	for i, planet := range g.Planets {
		if planet.Name == name {
			return i
		}
	}
	return -1
}

func (g *Game) getRecepieByName(name string) *Recepie {
	for _, recepie := range g.Recepies {
		if recepie.Result.getName() == name {
			return recepie
		}
	}
	return nil
}

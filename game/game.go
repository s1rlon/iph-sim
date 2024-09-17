package game

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Game struct {
	Planets   []*Planet
	LastSteps int
	Managers  []*Manager
	Projects  *Projects
	db        *sql.DB
	GamdeData *GameData
	Ships     *Ships
	Ores      []*Ore
	Recepies  []*Recepie
	Alloys    []*Alloy
	Items     []*Item
	Rooms     *Rooms
	Beacon    *Beacon
}

var GlobalCalcer *Calcer
var DB *sql.DB
var MarketSVC *Market

func NewGame() *Game {
	db, err := sql.Open("sqlite3", "ipm.sql")
	if err != nil {
		panic(err)
	}

	gameData := &GameData{UpgradeHistory: []UpgradeHistory{}, Smelters: 1, Crafters: 1, ManagerSlots: 2}
	gameData.LoadUpgradeHistoryFromDB(db)

	makeTables(db)
	ores := createOres()
	alloys := createAlloys()
	items := createItems()

	managers := getManagersFromDB(db)
	projects := loadProjectsFromDB(db)
	ships := loadShipsFromDB(db)
	planets := makeNewPlanets(ores)
	rooms := loadRoomsFromDB(db)
	beacon := loadBeaconDataFromDB(db)

	return &Game{
		Planets:   planets,
		LastSteps: 0,
		db:        db,
		Managers:  managers,
		Projects:  projects,
		GamdeData: gameData,
		Ships:     ships,
		Ores:      ores,
		Alloys:    alloys,
		Items:     items,
		Rooms:     rooms,
		Beacon:    beacon,
	}
}

func (g *Game) InitData() {
	GlobalCalcer = NewCalcer(g)
	MarketSVC = NewMarket(g)
	g.Recepies = createRecepies(g)
	DB = g.db
	dbPlanets, _ := getPlanetsFromDB(g.db)
	for _, planet := range g.Planets {
		for _, dbPlanet := range dbPlanets {
			if planet.Name == dbPlanet.Name {
				planet.MiningLevel = dbPlanet.MiningLevel
				planet.ShipSpeedLeve1 = dbPlanet.ShipSpeedLeve1
				planet.ShipCargoLevel = dbPlanet.ShipCargoLevel
				planet.Locked = dbPlanet.Locked
				planet.ColonyLevel = dbPlanet.ColonyLevel
				planet.AlchemyLevel = dbPlanet.AlchemyLevel
			}
		}
	}
}

func (g *Game) ResetGalaxy() {
	g.ResetPlanets()
	g.ResetManagers()
	g.GamdeData.resetGameData()
	g.Projects = newProjects()
}

func (g *Game) ResetPlanets() {
	for _, planet := range g.Planets {
		planet.resetPlanet()
	}
}

func (g *Game) ResetManagers() {
	for _, manager := range g.Managers {
		manager.unassignManager()
	}
}

func (g *Game) moneySpent() float64 {
	total := 0.0
	for _, upgrade := range g.GamdeData.UpgradeHistory {
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

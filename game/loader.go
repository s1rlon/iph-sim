package game

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadJSON(filePath string, target interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("failed to decode JSON from %s: %w", filePath, err)
	}

	return nil
}

type JSONRecipe struct {
	Result string             `json:"Result"`
	Input  map[string]float64 `json:"Input"`
}

func (g *Game) loadOres() error {
	var ores []*Ore
	if err := loadJSON("data/ores.json", &ores); err != nil {
		return err
	}
	g.Ores = ores
	return nil
}

func (g *Game) loadPlanets() error {
	var planetData []struct {
		Name         string    `json:"Name"`
		BaseOreNames []string  `json:"BaseOreNames"`
		Distribution []float64 `json:"Distribution"`
		UnlockCost   int       `json:"UnlockCost"`
		Distance     float64   `json:"Distance"`
	}

	if err := loadJSON("data/planets.json", &planetData); err != nil {
		return err
	}

	g.Planets = make([]*Planet, len(planetData))
	for i, pd := range planetData {
		planetOres := getOres(g.Ores, pd.BaseOreNames...)
		g.Planets[i] = NewPlanet(pd.Name, planetOres, pd.BaseOreNames, pd.Distribution, pd.UnlockCost, pd.Distance)
	}

	return nil
}

func (g *Game) loadAlloys() error {
	var alloys []*Alloy
	if err := loadJSON("data/alloys.json", &alloys); err != nil {
		return err
	}
	g.Alloys = alloys
	return nil
}

func (g *Game) loadItems() error {
	var items []*Item
	if err := loadJSON("data/items.json", &items); err != nil {
		return err
	}
	g.Items = items
	return nil
}

func (g *Game) loadRecipes() error {
	var jsonRecipes []JSONRecipe
	if err := loadJSON("data/recipes.json", &jsonRecipes); err != nil {
		return err
	}

	recipes := make([]*Recepie, 0, len(jsonRecipes))
	for _, jr := range jsonRecipes {
		result := g.getCratablebyName(jr.Result)
		if result == nil {
			continue
		}

		inputs := make(map[Craftable]float64)
		for name, qty := range jr.Input {
			input := g.getCratablebyName(name)
			if input != nil {
				inputs[input] = qty
			}
		}

		recipes = append(recipes, &Recepie{
			Result: result,
			Input:  inputs,
		})
	}
	g.Recepies = recipes
	return nil
}

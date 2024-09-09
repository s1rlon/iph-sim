package game

import (
	"testing"
)

func TestGetMiningRate(t *testing.T) {
	planet := &Planet{}
	expectedRate := 0.25
	actualRate := planet.getMiningRate(1)
	if actualRate != expectedRate {
		t.Errorf("Expected mining rate %f, but got %f", expectedRate, actualRate)
	}
}

func TestMine(t *testing.T) {
	planet := &Planet{
		Ores:         []Ore{{Name: "Copper"}, {Name: "Iron"}},
		Distribution: []float64{0.6, 0.4},
		MiningLevel:  1,
		Locked:       false,
	}
	expectedMined := map[string]float64{"Copper": 0.15, "Iron": 0.1}
	actualMined := planet.Mine(1)
	for ore, expectedAmount := range expectedMined {
		if actualMined[ore] != expectedAmount {
			t.Errorf("Expected mined amount of %s: %f, but got %f", ore, expectedAmount, actualMined[ore])
		}
	}
}

func TestGetLevelUpgradeCost(t *testing.T) {
	planet := &Planet{UnlockCost: 1000}
	expectedCost := 50.0
	actualCost := planet.getLevelUpgradeCost(1)
	if actualCost != expectedCost {
		t.Errorf("Expected level upgrade cost %f, but got %f", expectedCost, actualCost)
	}
}

func TestGetUpgradeCost(t *testing.T) {
	planet := &Planet{
		UnlockCost:     1000,
		MiningLevel:    1,
		ShipSpeedLeve1: 1,
		ShipCargoLevel: 1,
		Locked:         false,
	}
	expectedCost := planet.getLevelUpgradeCost(1)
	actualCost := planet.getUpgradeCost()
	if actualCost != expectedCost {
		t.Errorf("Expected upgrade cost %f, but got %f", expectedCost, actualCost)
	}
}

func TestGetShippingVolume(t *testing.T) {
	planet := &Planet{
		ShipSpeedLeve1: 2,
		ShipCargoLevel: 2,
		Distance:       16,
	}
	expectedVolume := planet.getShipSpeed(2) * planet.getShipCargo(2) / 16.0
	actualVolume := planet.getShippingVolume()
	if actualVolume != expectedVolume {
		t.Errorf("Expected shipping volume %f, but got %f", expectedVolume, actualVolume)
	}
}

func TestIsCargoSufficent(t *testing.T) {
	planet := &Planet{
		MiningLevel:    25,
		ShipSpeedLeve1: 11,
		ShipCargoLevel: 11,
		Distance:       10,
	}
	if !planet.isCargoSufficent(25) {
		t.Errorf("Expected cargo to be sufficient for mining level 1")
	}
}

func TestUpgradeMining(t *testing.T) {
	planet := &Planet{
		MiningLevel:    1,
		ShipSpeedLeve1: 1,
		ShipCargoLevel: 1,
		Distance:       10,
	}
	planet.upgradeMining()
	if planet.MiningLevel != 2 {
		t.Errorf("Expected MiningLevel to be 2, but got %d", planet.MiningLevel)
	}
	if planet.ShipSpeedLeve1 != 1 || planet.ShipCargoLevel != 1 {
		t.Errorf("Expected ShipSpeedLeve1 and ShipCargoLevel to be 1, but got %d and %d", planet.ShipSpeedLeve1, planet.ShipCargoLevel)
	}
}

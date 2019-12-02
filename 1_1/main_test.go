package main

import "testing"

func TestCalculateFuelRequirements(t *testing.T) {
	if result := calculateFuelRequirements(12); result != 2 {
		t.Errorf("Incorrect fuel requirement\nExpected %d\nGot%d\n", 2, result)
	}

	if result := calculateFuelRequirements(14); result != 2 {
		t.Errorf("Incorrect fuel requirement\nExpected %d\nGot%d\n", 2, result)
	}

	if result := calculateFuelRequirements(1969); result != 654 {
		t.Errorf("Incorrect fuel requirement\nExpected %d\nGot%d\n", 654, result)
	}

	if result := calculateFuelRequirements(100756); result != 33583 {
		t.Errorf("Incorrect fuel requirement\nExpected %d\nGot%d\n", 33583, result)
	}
}

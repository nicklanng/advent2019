package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	var fuelTotal int

	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to load input: %v\n", err)
	}

	r := bufio.NewReader(bytes.NewReader(b))

	for {
		mass, err := readNextModuleMass(r)
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		fuel := calculateFuelRequirements(mass)

		// calculate fuel required to carry fuel
		extraFuel := fuel
		for {
			extraFuel = calculateFuelRequirements(extraFuel)
			if extraFuel == 0 {
				break
			}
			fuel += extraFuel
		}

		log.Printf("Found module with mass: %d\tRequired fuel: %d\n", mass, fuel)
		fuelTotal += fuel
	}

	log.Printf("Total fuel needed: %d\n", fuelTotal)
}

func readNextModuleMass(r *bufio.Reader) (int, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}

	n, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		return 0, err
	}

	return n, nil
}

func calculateFuelRequirements(mass int) int {
	fuel := mass/3 - 2

	if fuel <0 {
		return 0
	}

	return fuel
}

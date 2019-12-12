package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type Vector3 struct {
	X int
	Y int
	Z int
}

func (v Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to load Input: %v\n", err)
	}

	var positions [4]Vector3
	var velocities [4]Vector3

	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.Trim(line, "<>")
		tokens := strings.Split(line, ", ")

		x, _ := strconv.Atoi(tokens[0][2:])
		y, _ := strconv.Atoi(tokens[1][2:])
		z, _ := strconv.Atoi(tokens[2][2:])

		positions[i] = Vector3{x, y, z}
	}

	fmt.Printf("After %d steps:\n", 0)
	for i := 0; i < 4; i++ {
		fmt.Println("pos=", positions[i], "vel=", velocities[i])
	}
	fmt.Println()

	for iteration := 0; iteration < 1000; iteration++ {
		for i := 0; i < 4; i++ {
			for j := i + 1; j < 4; j++ {
				iV := positions[i]
				jV := positions[j]

				if iV.X > jV.X {
					velocities[i].X--
					velocities[j].X++
				} else if iV.X < jV.X {
					velocities[i].X++
					velocities[j].X--
				}

				if iV.Y > jV.Y {
					velocities[i].Y--
					velocities[j].Y++
				} else if iV.Y < jV.Y {
					velocities[i].Y++
					velocities[j].Y--
				}

				if iV.Z > jV.Z {
					velocities[i].Z--
					velocities[j].Z++
				} else if iV.Z < jV.Z {
					velocities[i].Z++
					velocities[j].Z--
				}
			}
		}

		for i := 0; i < 4; i++ {
			positions[i] = positions[i].Add(velocities[i])
		}

		fmt.Printf("After %d steps:\n", iteration+1)
		for i := 0; i < 4; i++ {
			fmt.Println("pos=", positions[i], "vel=", velocities[i])
		}
		fmt.Println()
	}

	var total int

	for i := 0; i < 4; i++ {
		potentialEnergy := 0
		kineticEnergy := 0

		potentialEnergy += int(math.Abs(float64(positions[i].X)))
		potentialEnergy += int(math.Abs(float64(positions[i].Y)))
		potentialEnergy += int(math.Abs(float64(positions[i].Z)))

		kineticEnergy += int(math.Abs(float64(velocities[i].X)))
		kineticEnergy += int(math.Abs(float64(velocities[i].Y)))
		kineticEnergy += int(math.Abs(float64(velocities[i].Z)))

		total += potentialEnergy * kineticEnergy
	}

	fmt.Println(total)
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

	savedXStates := map[[8]int]int{}

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

	for iteration := 0; iteration < 186500; iteration++ {
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

		xState := [8]int{positions[0].X, positions[1].X, positions[2].X, positions[3].X, velocities[0].X, velocities[1].X, velocities[2].X, velocities[3].X}
		it, ok := savedXStates[xState]
		if !ok {
			savedXStates[xState] = iteration
		} else {
			fmt.Println(iteration+1, it, iteration-it)
		}
	}

	/*
		186028 231614 102356
		1, 2, 4, 46507, 93014, 186028
		1, 2, 115807, 231614
		1, 2, 4, 25589, 51178, 102356
	*/

	fmt.Println(93014 * 51178 * 115807)
}

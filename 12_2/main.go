package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Vector3 struct {
	X int16
	Y int16
	Z int16
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

	var vectors [8]Vector3
	states := map[[8]Vector3]struct{}{}

	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.Trim(line, "<>")
		tokens := strings.Split(line, ", ")

		x, _ := strconv.Atoi(tokens[0][2:])
		y, _ := strconv.Atoi(tokens[1][2:])
		z, _ := strconv.Atoi(tokens[2][2:])

		vectors[i] = Vector3{int16(x), int16(y), int16(z)}
	}

	fmt.Println(0, vectors)

	for iteration := 0; ; iteration++ {
		for i := 0; i < 4; i++ {
			for j := i + 1; j < 4; j++ {
				iV := vectors[i]
				jV := vectors[j]

				if iV.X > jV.X {
					vectors[i+4].X--
					vectors[j+4].X++
				} else if iV.X < jV.X {
					vectors[i+4].X++
					vectors[j+4].X--
				}

				if iV.Y > jV.Y {
					vectors[i+4].Y--
					vectors[j+4].Y++
				} else if iV.Y < jV.Y {
					vectors[i+4].Y++
					vectors[j+4].Y--
				}

				if iV.Z > jV.Z {
					vectors[i+4].Z--
					vectors[j+4].Z++
				} else if iV.Z < jV.Z {
					vectors[i+4].Z++
					vectors[j+4].Z--
				}
			}
		}

		for i := 0; i < 4; i++ {
			vectors[i] = vectors[i].Add(vectors[i+4])
		}

		// fmt.Println(iteration+1, vectors)

		_, ok := states[vectors]
		if ok {
			fmt.Printf("Duplicate found in step %d\n", iteration)
			break
		} else {
			states[vectors] = struct{}{}
		}
	}
}

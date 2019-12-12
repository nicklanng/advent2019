package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Vector3 struct {
	X int8
	Y int8
	Z int8
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
	states := map[[20]byte]struct{}{}

	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.Trim(line, "<>")
		tokens := strings.Split(line, ", ")

		x, _ := strconv.Atoi(tokens[0][2:])
		y, _ := strconv.Atoi(tokens[1][2:])
		z, _ := strconv.Atoi(tokens[2][2:])

		vectors[i*2] = Vector3{int8(x), int8(y), int8(z)}
	}

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

		state := sha1.Sum([]byte{byte(vectors[0].X), byte(vectors[0].Y), byte(vectors[0].Z), byte(vectors[1].X), byte(vectors[1].Y), byte(vectors[1].Z), byte(vectors[2].X), byte(vectors[2].Y), byte(vectors[2].Z), byte(vectors[3].X), byte(vectors[3].Y), byte(vectors[3].Z)})
		_, ok := states[state]
		if ok {
			fmt.Printf("Duplicate found in step %d\n", iteration)
			break
		} else {
			states[state] = struct{}{}
		}
	}
}

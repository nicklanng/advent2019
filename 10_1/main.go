package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

type Vector2 struct {
	X float64
	Y float64
}

func (v Vector2) Sub(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func (v Vector2) Normalize() Vector2{
	length := v.Magnitude()
	return Vector2{
		X: v.X / length,
		Y: v.Y / length,
	}
}

func (v Vector2) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v Vector2) EqualEpsilon(v2 Vector2, epsilon float64) bool {
	return math.Abs(v.X - v2.X) <= epsilon && math.Abs(v.Y - v2.Y) <= epsilon
}

func main() {
	asteroids := parseAsteroidsFile()
	fmt.Println(len(asteroids))

	mostSeenAsteroids := 0
	for _, origin := range asteroids {
		thisCount := 0

		for _, target := range asteroids {
			if origin == target {
				continue
			}

			blockerFound := false
			for _, blocker := range asteroids {
				if origin == blocker {
					continue
				}
				if target == blocker {
					continue
				}

				blockerFound = blocks(origin, target, blocker)
				if blockerFound {
					break
				}
			}
			if !blockerFound {
				thisCount++
			}
		}

		if thisCount > mostSeenAsteroids {
			mostSeenAsteroids = thisCount
		}
	}

	fmt.Println(mostSeenAsteroids)
}

func blocks(origin Vector2, target Vector2, blocker Vector2) bool {
	originToTargetDirection := target.Sub(origin).Normalize()
	originToBlockerDirection := blocker.Sub(origin).Normalize()

	// if vectors arent aligned, no dice
	if !originToTargetDirection.EqualEpsilon(originToBlockerDirection, .001) {
		return false
	}

	// if blocker is begind target, no block
	if blocker.Sub(origin).Magnitude() < target.Sub(origin).Magnitude() {
		return false
	}

	return true
}

func parseAsteroidsFile() []Vector2 {
	var asteroids []Vector2
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to load Input: %v\n", err)
	}


	x := 0
	y := 0
	for i := range b {
		switch b[i] {
		case '#':
			asteroids = append(asteroids, Vector2{float64(x), float64(y)})
		case '\n':
			x = 0
			y++
		}
		x++
	}

	return asteroids
}

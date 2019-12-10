package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
)

var up = Vector2{0, -1}
var origin = Vector2{27, 19}

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

type Asteroids []Vector2
func (a Asteroids) Len() int      { return len(a) }
func (a Asteroids) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Asteroids) Less(i, j int) bool {
	direction1 := a[i].Sub(origin)
	angle1 := math.Atan2(direction1.Y, direction1.X) - math.Atan2(up.Y, up.X)
	if angle1 < 0 {
		angle1 += 2 * math.Pi
	}

	direction2 := a[j].Sub(origin)
	angle2 := math.Atan2(direction2.Y, direction2.X) - math.Atan2(up.Y, up.X)
	if angle2 < 0 {
		angle2 += 2 * math.Pi
	}

	fmt.Println(angle1, angle2)

	return angle1 < angle2
}

func main() {
	asteroids := parseAsteroidsFile()
	fmt.Println(len(asteroids))

	var seen []Vector2

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
			seen = append(seen, target)
		}
	}

	sort.Sort(Asteroids(seen))
	for i :=0; i < 200; i++ {
		fmt.Println(i + 1, seen[i])
	}
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
		return true
	}

	return false
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
			x++
			break
		case '.':
			x++
			break
		case '\n':
			x = 0
			y++
			break
		}
	}

	return asteroids
}
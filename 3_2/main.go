package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type Orientation byte

const (
	Horizontal Orientation = 0
	Vertical Orientation = 1
)

type Vector2 struct {
	X int
	Y int
}

type Line struct {
	Start Vector2
	End Vector2
}

func (l Line) Orientation() Orientation {
	if l.Start.X != l.End.X {
		return Horizontal
	}
	return Vertical
}

func (l Line) Intersects(l2 Line) bool {
	if l.Orientation() == l2.Orientation() {
		return false
	}

	l = l.PositiveDirection()
	l2 = l2.PositiveDirection()

	if l.Orientation() == Horizontal {
		// line 1 is left of line 2
		if l.Start.X < l2.Start.X && l.End.X < l2.Start.X {
			return false
		}

		// line 2 is right of line 2
		if l.Start.X > l2.Start.X && l.End.X > l2.Start.X {
			return false
		}

		//is line above or below line 2
		if l.Start.Y < l2.Start.Y || l.Start.Y > l2.End.Y {
			return false
		}

		return true

	} else {
		// line 1 is below line 2
		if l.Start.Y < l2.Start.Y && l.End.Y < l2.Start.Y {
			return false
		}

		// line 2 is above line 2
		if l.Start.Y > l2.Start.Y && l.End.Y > l2.Start.Y {
			return false
		}

		// is line to left or right of line 2
		if l.Start.X < l2.Start.X || l.Start.X > l2.End.X {
			return false
		}

		return true
	}
}

func (l Line) PositiveDirection() Line {
	if l.Orientation() == Horizontal && l.Start.X > l.End.X {
		l.Start, l.End = l.End, l.Start
	}

	if l.Orientation() == Vertical && l.Start.Y > l.End.Y {
		l.Start, l.End = l.End, l.Start
	}

	return l
}

func (l Line) IntersectsPoint(p Vector2) bool {
	l = l.PositiveDirection()

	if l.Orientation() == Horizontal && l.Start.Y != p.Y {
		return false
	}

	if l.Orientation() == Vertical && l.Start.X != p.X {
		return false
	}

	return p.X >= l.Start.X && p.X <= l.End.X && p.Y >= l.Start.Y && p.Y <= l.End.Y
}

func (l Line) Length() int {
	return int(math.Abs(float64(l.Start.X-l.End.X)) + math.Abs(float64(l.Start.Y-l.End.Y)))
}

type Path struct {
	Lines []Line
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to load input: %v\n", err)
	}

	inputRows := strings.Split(string(b), "\n")
	paths := parsePaths(inputRows)
	collisions := PlotAndMeasure(paths[0], paths[1])

	bestDist := math.MaxInt64
	for _, c := range collisions {
		dist := 0
		for _, p := range paths {
			dist += countPathLength(p, c)
		}
		if dist < bestDist {
			bestDist = dist
		}
	}
	fmt.Println(bestDist)
}

func countPathLength(p Path, c Vector2) int {
	dist := 0
	for _, l := range p.Lines {
		if l.IntersectsPoint(c) {
			l.End = c
			dist += l.Length()
			return dist
		}

		dist += l.Length()
	}
	return dist
}

func parsePaths(inputRows []string) []Path {
	var result []Path

	for i := range inputRows {
		var currentPosition Vector2
		var path Path


		if len(inputRows[i]) == 0 {
			continue
		}

		tokens := strings.Split(inputRows[i], ",")
		for _, t := range tokens {

			command := t[0]
			dist, err := strconv.Atoi(t[1:])
			if err != nil {
				panic(err)
			}

			l := Line{Start: currentPosition}

			switch command {
			case 'U':
				currentPosition.Y += dist
				break
			case 'D':
				currentPosition.Y -= dist
				break
			case 'R':
				currentPosition.X += dist
				break
			case 'L':
				currentPosition.X -= dist
				break
			}

			l.End = currentPosition

			path.Lines = append(path.Lines, l)
		}

		result = append(result, path)
	}

	return result
}

func PlotAndMeasure(path1, path2 Path) []Vector2 {
	var result []Vector2

	for _, l1 := range path1.Lines {
		for _, l2 := range path2.Lines {
			if l1.Intersects(l2) {
				if l1.Orientation() == Horizontal {
					result = append(result, Vector2{
						X: l2.Start.X,
						Y: l1.Start.Y,
					})
				} else {
					result = append(result, Vector2{
						X: l1.Start.X,
						Y: l2.Start.Y,
					})
				}
			}
		}
	}

	return result
}
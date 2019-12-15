package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"time"
)

type Direction byte

const (
	North Direction = iota + 1
	South
	West
	East
)

type Robot struct {
	Position Vector2
	Computer *Computer
	Floor    map[Vector2]byte

	OxygenLocation Vector2
}

func (r *Robot) Run() {
	done := make(chan struct{})

	go func() {
		err := r.Computer.Run()
		done <- struct{}{}
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from", r)
			}
		}()

		r.Position = Vector2{25, 25}
		r.Floor[r.Position] = '.'
		r.search(North, South)
		r.search(South, North)
		r.search(West, East)
		r.search(East, West)

		minutes := 0
		frontier := []Vector2{r.OxygenLocation}
		r.Floor[r.OxygenLocation] = 'g'

		for {
			time.Sleep(time.Second / 10)

			img := image.NewRGBA(image.Rectangle{Max: image.Point{50, 50}})
			for pos, c := range r.Floor {
				switch c {
				case '#':
					img.Set(pos.X, pos.Y, color.White)
				case '.':
					img.Set(pos.X, pos.Y, color.Black)
				case '*':
					img.Set(pos.X, pos.Y, color.RGBA{0x0, 0x0, 0xff, 0xff})
				case 'o':
					img.Set(pos.X, pos.Y, color.RGBA{0x0, 0xff, 0x0, 0xff})
				case 'g':
					img.Set(pos.X, pos.Y, color.RGBA{0xff, 0x0, 0xff, 0xff})
				}
			}

			f, _ := os.Create("out.png")
			if err := png.Encode(f, img); err != nil {
				panic(err)
			}

			newFrontier := []Vector2{}
			spread := false
			for _, v := range frontier {
				n := r.GetNewPosition(v, North)
				if r.Floor[n] == '.' {
					newFrontier = append(newFrontier, n)
					r.Floor[n] = 'g'
					spread = true
				}
				s := r.GetNewPosition(v, South)
				if r.Floor[s] == '.' {
					newFrontier = append(newFrontier, s)
					r.Floor[s] = 'g'
					spread = true
				}
				w := r.GetNewPosition(v, West)
				if r.Floor[w] == '.' {
					newFrontier = append(newFrontier, w)
					r.Floor[w] = 'g'
					spread = true
				}
				e := r.GetNewPosition(v, East)
				if r.Floor[e] == '.' {
					newFrontier = append(newFrontier, e)
					r.Floor[e] = 'g'
					spread = true
				}
			}

			frontier = newFrontier

			if spread {
				minutes++
				spread = false
			}

			if len(frontier) == 0 {
				break
			}
		}

		fmt.Println(minutes)
		done <- struct{}{}
	}()

	<-done
}

func (r *Robot) search(d Direction, back Direction) {
	v := r.GetNewPosition(r.Position, d)

	// skip already searched
	if _, ok := r.Floor[v]; ok {
		return
	}

	r.Computer.Input <- fmt.Sprintf("%d", d)
	val, _ := strconv.Atoi(<-r.Computer.Output)

	switch val {
	case 0:
		r.Floor[v] = '#'
		return
	case 1:
		r.Floor[v] = '.'
		r.Position = v
		r.search(North, South)
		r.search(South, North)
		r.search(West, East)
		r.search(East, West)
	case 2:
		r.Floor[v] = 'o'
		r.Position = v
		r.OxygenLocation = v
		r.search(North, South)
		r.search(South, North)
		r.search(West, East)
		r.search(East, West)
	}

	// go back
	r.Computer.Input <- fmt.Sprintf("%d", back)
	<-r.Computer.Output
	r.Position = r.GetNewPosition(r.Position, back)
}

func (r *Robot) GetNewPosition(v Vector2, d Direction) Vector2 {
	switch d {
	case North:
		return Vector2{v.X, v.Y - 1}
	case South:
		return Vector2{v.X, v.Y + 1}
	case West:
		return Vector2{v.X - 1, v.Y}
	case East:
		return Vector2{v.X + 1, v.Y}
	}
	return Vector2{}
}

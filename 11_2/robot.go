package main

import (
	"fmt"
	"strconv"
)

type Direction byte

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Robot struct {
	Direction Direction
	Position Vector2
	Computer *Computer
	Floor map[Vector2]byte

	LowestX int
	LowestY int
	LargestX int
	LargestY int

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

		for {
			// read color
			floorColor := r.Floor[r.Position]
			r.Computer.Input <- fmt.Sprintf("%d", floorColor)

			// paint color
			colorToPaint, _ := strconv.ParseUint(<-r.Computer.Output, 10, 8)
			color := byte(colorToPaint)
			r.Floor[r.Position] = color

			// turn
			direction, _ := strconv.ParseInt(<-r.Computer.Output, 10, 64)
			switch direction {
			case 0:
				r.Direction = (r.Direction - 1) % 4
				break
			case 1:
				r.Direction = (r.Direction + 1) % 4
				break
			}

			// move forward
			switch r.Direction {
			case Up:
				r.Position = Vector2{r.Position.X, r.Position.Y-1}
				break
			case Down:
				r.Position = Vector2{r.Position.X, r.Position.Y+1}
				break
			case Left:
				r.Position = Vector2{r.Position.X-1, r.Position.Y}
				break
			case Right:
				r.Position = Vector2{r.Position.X+1, r.Position.Y}
				break
			}

			if r.Position.X < r.LowestX {
				r.LowestX = r.Position.X
			}

			if r.Position.Y < r.LowestY {
				r.LowestY = r.Position.Y
			}

			if r.Position.X > r.LargestX {
				r.LargestX = r.Position.X
			}

			if r.Position.Y > r.LargestY {
				r.LargestY = r.Position.Y
			}
		}

	}()

	<- done
}
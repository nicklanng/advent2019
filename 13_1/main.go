package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Arcade struct {
	computer *Computer
	display  *image.RGBA
}

func (a *Arcade) Run() {
	go func() {
		if err := a.computer.Run(); err != nil {
			panic(err)
		}
	}()

	for {
		sx, ok := <-a.computer.Output
		if !ok {
			break
		}
		sy, ok := <-a.computer.Output
		if !ok {
			break
		}
		stile, ok := <-a.computer.Output
		if !ok {
			break
		}

		x, _ := strconv.Atoi(sx)
		y, _ := strconv.Atoi(sy)
		tile, _ := strconv.Atoi(stile)

		switch tile {
		case 0:
			a.display.Set(x, y, color.RGBA{0x0, 0x0, 0x0, 0xff})
			break
		case 1:
			a.display.Set(x, y, color.RGBA{0xe2, 0xf3, 0xe4, 0xff})
			break
		case 2:
			a.display.Set(x, y, color.RGBA{0x33, 0x2c, 0x50, 0xff})
			break
		case 3:
			a.display.Set(x, y, color.RGBA{0x46, 0x87, 0x8f, 0xff})
			break
		case 4:
			a.display.Set(x, y, color.RGBA{0x94, 0xe3, 0x44, 0xff})
			break
		}
	}

}

func main() {
	program := loadProgramFromDisk()

	c := &Computer{}
	c.Load(program)

	a := &Arcade{
		computer: c,
		display:  image.NewRGBA(image.Rectangle{Max: image.Point{X: 43, Y: 21}}),
	}

	a.Run()

	f, _ := os.Create("out.png")
	if err := png.Encode(f, a.display); err != nil {
		panic(err)
	}

	// count blocks
	var i int
	for x := 0; x < 43; x++ {
		for y := 0; y < 21; y++ {
			if a.display.At(x, y) == (color.RGBA{0x33, 0x2c, 0x50, 0xff}) {
				i++
			}
		}
	}
	fmt.Println(i)
}

func loadProgramFromDisk() []int {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to load Input: %v\n", err)
	}
	text := strings.Split(string(b), ",")
	program := make([]int, len(text))
	for i := range text {
		val, err := strconv.Atoi(text[i])
		if err != nil {
			panic(err)
		}
		program[i] = val
	}
	return program
}

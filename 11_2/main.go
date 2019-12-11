package main

import (
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	program := loadProgramFromDisk()

	c := &Computer{}
	c.Load(program)

	floor := map[Vector2]byte{}
	floor[Vector2{}] = 1

	r := &Robot{
		Computer:  c,
		Floor:     floor,
	}

	r.Run()

	img := image.NewRGBA(image.Rectangle{Max: image.Point{X: r.LargestX - r.LowestY + 1, Y: r.LargestY - r.LowestY + 1}})
	for pos, c := range floor {
		switch c {
		case 0:
			img.Set(pos.X - r.LowestX, pos.Y - r.LowestY, color.Black)
		case 1:
			img.Set(pos.X - r.LowestX, pos.Y - r.LowestY, color.White)
		}
	}

	f, _ := os.Create("out.png")
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
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

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	program := loadProgramFromDisk()

	c := &Computer{}
	c.Load(program)

	go c.Run()

	positions := make(map[Vector2]struct{})

	y := 0
	x := 0
	for out := range c.Output {
		char, _ := strconv.Atoi(out)
		if char == 10 {
			y++
			x = 0
			fmt.Printf("%c", char)
			continue
		}
		if char == 35 {
			positions[Vector2{x, y}] = struct{}{}
		}

		fmt.Printf("%c", char)
		x++
	}

	sum := 0
	for k := range positions {
		_, n := positions[Vector2{k.X, k.Y - 1}]
		_, s := positions[Vector2{k.X, k.Y + 1}]
		_, w := positions[Vector2{k.X - 1, k.Y}]
		_, e := positions[Vector2{k.X + 1, k.Y}]

		if n && s && w && e {
			sum += k.X * k.Y
		}
	}
	fmt.Println(sum)
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

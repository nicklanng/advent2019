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

	floor := map[Vector2]byte{}

	r := &Robot{
		Computer:  c,
		Floor:     floor,
	}

	r.Run()

	fmt.Println(len(floor))
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

package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to load input: %v\n", err)
	}

	text := strings.Split(string(b), ",")
	program := make([]int, len(text))
	for i := range text{
		val, err := strconv.Atoi(text[i])
		if err != nil {
			panic(err)
		}
		program[i] =val
	}

	computer := &Computer{}
	computer.Load(program)
	if err := computer.Run(); err != nil {
		log.Println(err)
	}
}
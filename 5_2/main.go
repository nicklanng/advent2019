package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to load input: %v\n", err)
	}

	program := strings.Split(string(b), ",")

	computer := &Computer{}
	computer.Load(program)
	if err := computer.Run(); err != nil {
		log.Println(err)
	}
}
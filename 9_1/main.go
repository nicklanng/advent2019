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

	done := make(chan struct{})

	go func() {
		for msg := range c.Output {
			fmt.Println(msg)
		}
		done <- struct{}{}
	}()

	go func() {
		for {
			var val string
			if _, err := fmt.Scan(&val); err != nil {
				panic(err)
			}
			c.Input <- val
		}
	}()

	if err := c.Run(); err != nil {
		panic(err)
	}

	<-done
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

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"sync"
)

func main() {
	program := loadProgramFromDisk()

	highestSignal := 0

	var perms []string
	perm([]rune("56789"), &perms, 0)

	for i := range perms {
		permutation := perms[i]
		fmt.Println("Permutation: " + permutation)

		var wg sync.WaitGroup
		wg.Add(5)
		var computers []*Computer
		for j := 0; j < 5; j++{
			c := &Computer{}
			c.Load(program)
			go func() {
				if err := c.Run(); err != nil {
					log.Println(err)
				}
				wg.Done()
			}()
			computers = append(computers, c)
		}

		go connect(computers[1].Input, computers[0].Output)
		go connect(computers[2].Input, computers[1].Output)
		go connect(computers[3].Input, computers[2].Output)
		go connect(computers[4].Input, computers[3].Output)
		go func() {
			result := connect(computers[0].Input, computers[4].Output)
			signal, _ := strconv.Atoi(result)
			if signal > highestSignal {
				highestSignal = signal
			}
		}()

		for j, c := range computers {
			c.Input <- string(permutation[j])
		}
		computers[0].Input <- "0"

		wg.Wait()
	}

	fmt.Println(highestSignal)
}

func connect(in, out chan string) (val string) {
	// hacky way to recover from write to closed channel
	defer func() {
		recover()
	}()

	for val = range out {
		in <- val
	}

	return
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

func perm(a []rune, s *[]string, i int) {
	if i > len(a) {
		*s = append(*s, string(a))
		return
	}
	perm(a, s, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, s, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
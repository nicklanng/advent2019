package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	thrusterIDs = "ABCDE"
)

func main() {
	program := loadProgramFromDisk()

	highestSignal := 0
	computer := NewComputer()

	var perms []string
	perm([]rune("01234"), &perms, 0)

	for i := range perms {
		permutation := perms[i]
		fmt.Println("Permutation: " + permutation)
		input := "0"
		for i := range permutation {
			computer.Load(program)
			go func() {
				if err := computer.Run(); err != nil {
					log.Println(err)
				}
			}()
			fmt.Printf("Thruster ID %c has input: %s\n", thrusterIDs[i], input)
			computer.Input <- string(permutation[i])
			computer.Input <- input
			input = <-computer.Output
			fmt.Printf("Thruster ID %c has output: %s\n", thrusterIDs[i], input)
		}

		signal, _ := strconv.Atoi(input)

		if signal > highestSignal {
			highestSignal = signal
		}
	}


	fmt.Println(highestSignal)
}

func loadProgramFromDisk() []int {
	b, err := ioutil.ReadFile("Input.txt")
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
		fmt.Println(a)
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
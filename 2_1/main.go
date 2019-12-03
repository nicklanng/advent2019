package main

import (
	"fmt"
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

	input := []int{}
	for _, str := range strings.Split(string(b), ",") {
		i, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			log.Fatalf("Invalid entry: %s", str)
		}
		input = append(input, i)
	}

	// restore gravity to sleighmobile 5000
	input[1] = 12
	input[2] = 2

	if err := process(input); err != nil {
		log.Fatalf("Failed to process input: %v", err)
	}

	log.Println(input[0])
}

func process(input []int) error {

	for i := 0; i < len(input); i += 4 {
		switch input[i] {
		case 1:
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
			break
		case 2:
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
			break
		case 99:
			return nil
		default:
			return fmt.Errorf("unknown opcode %d", input[i])
		}
	}

	return nil
}
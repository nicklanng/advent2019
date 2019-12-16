package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

var basePattern = []int{0, 1, 0, -1}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to load Input: %v\n", err)
	}

	numbers := make([]int, len(b))
	for i := range b {
		numbers[i] = int(b[i] - '0')
	}

	// run fft
	for i := 0; i < 100; i++ {
		numbers = fft(numbers)
		fmt.Println(numbers[:8])
	}

}

func fft(numbers []int) []int {
	output := make([]int, len(numbers))

	for i := range numbers {
		pattern := makePattern(i + 1)
		output[i] = sum(numbers, pattern)
		output[i] = output[i] % 10
	}

	return output
}

func makePattern(iteration int) []int {
	pattern := make([]int, len(basePattern)*iteration)

	for i := 0; i < len(basePattern)*iteration; i++ {
		pattern[i] = basePattern[(i/iteration)%4]
	}

	return pattern
}

func sum(numbers, pattern []int) int {
	sum := 0

	for i := range numbers {
		sum += numbers[i] * pattern[((i+1)%len(pattern))]
	}

	return int(math.Abs(float64(sum)))
}

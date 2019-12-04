package main

import "fmt"

//--- Day 4: Secure Container ---
//You arrive at the Venus fuel depot only to discover it's protected by a password. The Elves had written the password on a sticky note, but someone threw it out.
//
//However, they do remember a few key facts about the password:
//
//It is a six-digit number.
//The value is within the range given in your puzzle input.
//Two adjacent digits are the same (like 22 in 122345).
//Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
//Other than the range rule, the following are true:
//
//111111 meets these criteria (double 11, never decreases).
//223450 does not meet these criteria (decreasing pair of digits 50).
//123789 does not meet these criteria (no double).
//How many different passwords within the range given in your puzzle input meet these criteria?
//
//Your puzzle input is 134792-675810.

func main() {
	var possibleSolutions int

	for i := 134792; i < 675810; i++ {
		if !increasingNumbers(i) {
			continue
		}

		if !sameAdjacentDigits(i) {
			continue
		}

		fmt.Println(i)

		possibleSolutions++
	}

	fmt.Println(possibleSolutions)
}

func increasingNumbers(i int) bool {
	str := fmt.Sprintf("%d", i)

	for i := 1; i < len(str); i++ {
		if str[i] < str[i-1] {
			return false
		}
	}

	return true
}

func sameAdjacentDigits(i int) bool {
	str := fmt.Sprintf("%d", i)

	currentDigit := byte(0)
	count := 0

	for i := 0; i < len(str); i++ {
		if str[i] == currentDigit {
			count++
		} else {
			if count == 2 {
				return true
			}

			currentDigit = str[i]
			count = 1
		}
	}

	if count == 2 {
		return true
	}

	return false
}
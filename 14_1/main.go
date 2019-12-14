package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type ChemicalQuanity struct {
	Chemical string
	Quanity  int
}

type Recipe func() (quantityCreated int, inputs []ChemicalQuanity)

var (
	recipes = make(map[string]Recipe)
)

func parseChemicalQuantity(s string) ChemicalQuanity {
	tokens := strings.Split(s, " ")
	quanity, _ := strconv.Atoi(tokens[0])
	return ChemicalQuanity{strings.TrimSpace(tokens[1]), quanity}
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to load Input: %v\n", err)
	}

	lines := strings.Split(string(b), "\n")
	for _, l := range lines {
		recipe := strings.Split(strings.TrimSpace(l), " => ")

		output := parseChemicalQuantity(recipe[1])

		i := []ChemicalQuanity{}
		for _, in := range strings.Split(recipe[0], ", ") {
			input := parseChemicalQuantity(in)
			i = append(i, input)
		}

		recipes[output.Chemical] = func() (int, []ChemicalQuanity) {
			return output.Quanity, i
		}
	}

	materials := make(map[string]int)
	materials["FUEL"] = 1

	done := false
	for {
		if done {
			fmt.Println(materials["ORE"])
			break
		}

		done = true

		for k, v := range materials {

			// ore has no recipe
			if k == "ORE" {
				continue
			}

			// clean up materials that are not needed and have no surplus
			if materials[k] == 0 {
				delete(materials, k)
				continue
			}

			// we have a surplus of this material so skip
			if v <= 0 {
				continue
			}

			done = false

			// run recipe
			quanity, inputs := recipes[k]()
			materials[k] = v - quanity

			// add the required materials for the recipe
			for _, i := range inputs {
				materials[i.Chemical] = i.Quanity + materials[i.Chemical]
			}
		}
	}

}

// 10 ORE => 10 A
// 1 ORE => 1 B
// 7 A, 1 B => 1 C
// 7 A, 1 C => 1 D
// 7 A, 1 D => 1 E
// 7 A, 1 E => 1 FUEL

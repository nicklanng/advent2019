package main

import (
	"errors"
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

	materials["FUEL"] = 10000
	process(materials)

	chemicalsFast := []ChemicalQuanity{}
	for k, v := range materials {
		if v == 0 {
			continue
		}
		chemicalsFast = append(chemicalsFast, ChemicalQuanity{k, v})
	}

	recipes["FUELFAST"] = func() (int, []ChemicalQuanity) {
		return 1, chemicalsFast
	}
	fmt.Println(materials)

	fuel := 0
	materials = make(map[string]int)
	for {

		materials["FUELFAST"] = 1
		if err = process(materials); err != nil {
			break
		}
		fuel += 10000

		if materials["ORE"] > 975000000000 {
			break
		}
	}
	fmt.Println(materials)

	for {
		materials["FUEL"] = 1
		if err = process(materials); err != nil {
			break
		}
		fuel++
	}
	fmt.Println(materials)
	fmt.Println(fuel)

}

func process(materials map[string]int) error {
	done := false
	for {
		if done {
			break
		}

		done = true

		for k, v := range materials {

			// ore has no recipe
			if k == "ORE" {
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
				if i.Chemical == "ORE" {
					if materials["ORE"] > 1000000000000 {
						fmt.Println(materials["ORE"], i)
						return errors.New("out of ore")
					}
				}
				materials[i.Chemical] = i.Quanity + materials[i.Chemical]
			}
		}
	}

	return nil
}

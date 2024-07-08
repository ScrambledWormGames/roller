package main

import (
	"flag"
	"fmt"
	"math/rand"
	"slices"
)

func roll(die int, multi int, modifier int, breakdown bool) (result int) {
	// This func takes any multiplier of dice of a given type, rolls them, and then adds a modifier to the result
	rollTotal := 0

	fmt.Println("------------------------------------------------------------------")
	fmt.Printf("\tDie: %d * d%-6d | Modifier: %-3d | Breakdown: %t\n", multi, die, modifier, breakdown)
	fmt.Printf("##################################################################\n\n")

	for range multi {
		rollResult := 1 + rand.Intn(die)
		rollTotal += rollResult
		if breakdown {
			fmt.Printf("Rolled: %-3d\t\t Running total: %d\n", rollResult, rollTotal)
		}

	}

	if breakdown && modifier != 0 {
		fmt.Printf("Applying modifier %d\n\n", modifier)
	}

	return rollTotal + modifier
}

func main() {

	fmt.Println("##################################################################")
	fmt.Println("\t\t\tSUPER ROLLER v1.0")

	// Set out the dice we allow to be selected
	goodDice := []int{4, 6, 8, 10, 12, 20, 100}

	// Use flags to get the pointer of the die and parse out
	dSelect := flag.Int("d", 6, "Dice selection, out of: D4, D6, D8, D10, D12, D20 and D100")
	nSelect := flag.Int("n", 1, "The number of dice you want to roll")
	mSelect := flag.Int("m", 0, "Any modifier you wish to add")
	breakdown := flag.Bool("b", false, "Whether you want a breakdown of the dice rolls")
	flag.Parse()

	// Check for existence of the chosen die within our selection criteria
	if slices.Contains(goodDice, *dSelect) {
		fmt.Printf("Total: %d\n\n", roll(*dSelect, *nSelect, *mSelect, *breakdown))
	} else {
		fmt.Printf("##################################################################\n\n")
		fmt.Printf("Invalid dice d%d\n", *dSelect)
	}

}

// Package main proves Monty Hall's infamous problem once and for all.
//
// The problem is simple:
//
// Suppose you're on a game show, and you're given the choice of three doors:
// Behind one door is a car; behind the others, goats. You pick a door, say
// No. 1, and the host, who knows what's behind the doors, opens another door,
// say No. 3, which has a goat. He then says to you, "Do you want to pick door
// No. 2?" Is it to your advantage to switch your choice?
//
// For a more information, see Wikipedia's article on the [Monty Hall probelem].
//
// [Monty Hall problem]: https://en.wikipedia.org/wiki/Monty_Hall_problem
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Total number of people who try the problem.
	const people int = 1000
	// Number of people who who by sticking to their guess.
	var stuck int
	// Number of people who who by swaped to their guess.
	var swaped int
	// Split people into two catigories, those who stay with their inital guess,
	// and those who switch their guess after a useless door is revealed.
	for i := 0; i < people/2; i++ {
		stk, swp := runProblem()
		if stk {
			stuck++
		}
		if swp {
			swaped++
		}
	}
	fmt.Printf("%.2f%% of people who stuck with there guess won\n", (float32(stuck)/float32(people))*100)
	fmt.Printf("%.2f%% of people who swaped there guess won\n", (float32(swaped)/float32(people))*100)
}

// Function runProblem sets up and runs a random scenario.
func runProblem() (bool, bool) {

	// Hide prize
	var doors [3]bool
	doors[rand.Intn(len(doors))] = true

	// Pick door
	pick := rand.Intn(len(doors))

	// Reveal door with trash inside
	var trash []int
	for i, v := range doors {
		if v == false && i != pick {
			trash = append(trash, i)
		}
	}
	reveal := trash[rand.Intn(len(trash))]

	// Choose other door
	var other []int
	for i := range doors {
		if i != reveal && i != pick {
			other = append(other, i)
		}
	}
	swap := other[rand.Intn(len(other))]

	// Return result
	return doors[pick], doors[swap]

}

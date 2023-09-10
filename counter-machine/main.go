// Package CountMachine impelments an abstract Counter Machine used in
// theoretical computer science as a model of computation. A counter
// machine comprises a set of one or more unbounded registers, each of
// which can hold a single non-negative integer, and a list of arithmetic
// and control instructions for the machine to follow.
//
// See Wikipedia's article on [Counter Machines](https://en.wikipedia.org/wiki/Counter_machine)
// for a more detailed explaination.
//
// Inspired by Computerphile's video on [Counter Machines](https://www.youtube.com/watch?v=PXN7jTNGQIw).
package CountMachine

// Enabling Hardware Exceleration speeds up basic operations by making use of
// there physical implementations on clasical computers.
var HardwareExcel bool = false

// A Register can hold a single non-negative integer.
type Register uint64

type State interface {
	Run() int
}

// Check is a type used to preform checks on a register and continue
// to the next State acordingly.
type Check struct {
	Register *Register
	// The state to jump to when the register is Empty.
	Empty int
	// The state to jump to when the register is non-empty.
	NonEmpty int
}

func (c *Check) Run() int {
	if *c.Register <= 0 {
		return c.Empty
	} else {
		return c.NonEmpty
	}
}

// Action is a type used to preform an action on the register and then continue
// to the next State.
type Action struct {
	Register *Register
	// Set to true to increase the register and false to reduce the register.
	Sign bool
	// Move on the Next state when finished.
	Next int
}

func (a *Action) Run() int {
	if a.Sign {
		*a.Register++
	} else {
		*a.Register--
	}
	return a.Next
}

// Add is a type used to add to registers together.
type Add struct {
	C1   *Register
	C2   *Register
	Next int
}

func (a *Add) Run() int {
	if HardwareExcel {
		*a.C1 += *a.C2
		*a.C2 = 0
	} else {
		Exec(&[]State{
			&Entry{1},
			&Check{a.C2, 0, 2},
			&Action{a.C2, false, 3},
			&Action{a.C1, true, 1},
		})
	}
	return a.Next
}

// Entry is a type used like a pointer to the entrypoint State.
type Entry struct {
	Next int
}

func (h *Entry) Run() int {
	return h.Next
}

// Loop through a slice of States calling Run() on each untill it returns zero.
func Exec(states *[]State) {
	for state := (*states)[0].Run(); state != 0; state = (*states)[state].Run() {
	}
}

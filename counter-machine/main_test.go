package CountMachine

import (
	"testing"
)

func TestReset(t *testing.T) {
	var c1 Register = 13
	var states = []State{
		&Entry{1},
		&Check{&c1, 0, 2},
		&Action{&c1, false, 1},
	}
	Exec(&states)
	if c1 != 0 {
		t.Error(`c1 != 0`)
	}
}

func TestAdd(t *testing.T) {
	var c1, c2 Register = 13, 15
	var states = []State{
		&Entry{1},
		&Add{&c1, &c2, 0},
	}
	Exec(&states)
	if c1 != 28 {
		t.Errorf(`c1 = 28 got %d`, c1)
	}
}

func BenchmarkAdd(b *testing.B) {
	var c1, c2 Register
	var states = []State{
		&Entry{1},
		&Add{&c1, &c2, 0},
	}
	for i := 0; i < b.N; i++ {
		c2 = 100
		Exec(&states)
	}
}

func BenchmarkAddHardwareExcel(b *testing.B) {
	HardwareExcel = true
	defer func() { HardwareExcel = false }()

	var c1, c2 Register
	var states = []State{
		&Entry{1},
		&Add{&c1, &c2, 0},
	}
	for i := 0; i < b.N; i++ {
		c2 = 100
		Exec(&states)
	}
}

# CountMachine

Package CountMachine impelments an abstract Counter Machine used in theoretical computer science as a model of computation. A counter machine comprises a set of one or more unbounded registers, each of which can hold a single non-negative integer, and a list of arithmetic and control instructions for the machine to follow.

See Wikipedia's article on [Counter Machines](https://en.wikipedia.org/wiki/Counter_machine) for a more detailed explaination.

Inspired by Computerphile's video on [Counter Machines](https://www.youtube.com/watch?v=PXN7jTNGQIw).

## Invocation

You can instantiate Registers like this:

```go
var r1 Register = 13
```

Program states are repersented in memory as types that satisfy the `State` interface. Each state can contain different feilds based on the underling concrete type. References to other States in the same program are identifed by their index in the program slice. Registers are passed in via pointers. Here is a example program that resets `r1` to zero.

```go
var states = []State{
	&Entry{1},
	&Check{&r1, 0, 2},
	&Action{&r1, false, 1},
}
```

To run the program call `Exec()` on it.

```go
Exec(&states)
```

## Testing

```sh
go test
```

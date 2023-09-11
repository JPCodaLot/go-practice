// Package main implements basic buffer counters.
//
// This is an exercise project from [The Go Programming Language] book.
// The original example was called [bytecounter].
//
// [The Go Programming Language]: https://www.gopl.io/
// [bytecounter]: https://github.com/adonovan/gopl.io/blob/master/ch7/bytecounter/main.go
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	var count int
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	var count int
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		count++
	}
	*c += LineCounter(count)
	return count, nil
}

func main() {
	var b ByteCounter
	var w WordCounter
	var l LineCounter
	var name = "Dolly"
	fmt.Fprintf(&b, "Hello %s\nHow are you?", name)
	fmt.Fprintf(&w, "Hello %s\nHow are you?", name)
	fmt.Fprintf(&l, "Hello %s\nHow are you?", name)
	fmt.Println(b) // "24"
	fmt.Println(w) // "5"
	fmt.Println(l) // "2"
}

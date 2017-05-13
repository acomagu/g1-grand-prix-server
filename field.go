package main

import (
	"fmt"
	"os"
)

// The number of cell by one side.
const l = 13

// State is whether the cell is black, white or empty.
type State int

const (
	Empty State = iota
	Black
	White
)

// Field is whole game board.
type Field [l][l]State

func printField(field Field) {
	fmt.Fprint(os.Stderr, "  ")
	
	for x := 0; x < l; x++ {
		fmt.Fprintf(os.Stderr, "%2d", x)
	}
	fmt.Fprintln(os.Stderr, )

	for y := 0; y < l; y++ {
		fmt.Fprintf(os.Stderr, "%2d", y)
		for x := 0; x < l; x++ {
			switch field[y][x] {
			case Empty:
				fmt.Fprint(os.Stderr, "  ")
			case Black:
				fmt.Fprint(os.Stderr, " o")
			case White:
				fmt.Fprint(os.Stderr, " x")
			}
		}
		fmt.Fprintf(os.Stderr, "%2d\n", y)
	}

	fmt.Fprint(os.Stderr, "  ")
	for x := 0; x < l; x++ {
		fmt.Fprintf(os.Stderr, "%2d", x)
	}
	fmt.Fprintln(os.Stderr, )
}

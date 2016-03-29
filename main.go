package main

import (
	"bytes"
	"fmt"
)

// board size
const (
	SIZE = 3
)

type box int
type board [SIZE * SIZE]box
type boards []board

// box state
const (
	U box = iota
	X
	Y
)

func (b board) winner() bool {
	if b[0] != U {
		if b[0] == b[1] && b[1] == b[2] {
			return true
		}
		if b[0] == b[3] && b[3] == b[6] {
			return true
		}
		if b[0] == b[4] && b[4] == b[8] {
			return true
		}
	}

	if b[1] != U {
		if b[1] == b[4] && b[4] == b[7] {
			return true
		}
	}

	if b[2] != U {
		if b[2] == b[5] && b[5] == b[8] {
			return true
		}
		if b[2] == b[4] && b[4] == b[6] {
			return true
		}
	}

	if b[3] != U {
		if b[3] == b[4] && b[4] == b[5] {
			return true
		}
	}

	if b[6] != U {
		if b[6] == b[7] && b[7] == b[8] {
			return true
		}
	}

	return false
}

func (b board) complete() bool {
	for _, v := range b {
		if v == U {
			return false
		}
	}

	return true
}

func (b board) String() (res string) {
	var out bytes.Buffer

	for i := 0; i < len(b); i++ {
		if b[i] == X {
			out.WriteString("x")
		} else if b[i] == Y {
			out.WriteString("y")
		} else if b[i] == U {
			out.WriteString(" ")
		}

		if i+1 == SIZE*SIZE {
			out.WriteString("\n")
		} else if (i+1)%SIZE == 0 {
			out.WriteString("\n-----\n")
		} else {
			out.WriteString("|")
		}
	}

	return out.String()
}

func genBoardsR(b board, turn box) (res boards) {
	if b.complete() || b.winner() {
		res = append(res, b)
		return
	}

	for i := 0; i < len(b); i++ {
		if b[i] == U {
			bNew := b
			bNew[i] = turn

			next := X
			if turn == X {
				next = Y
			}

			bs := genBoardsR(bNew, next)
			res = append(res, bs...)
		}
	}

	return
}

func genBoards() (res boards) {
	xBoards := genBoardsR(board{}, X)
	res = xBoards

	//yBoards := genBoardsR(board{}, Y)
	//res = append(res, yBoards...)

	return
}

func main() {
	bs := genBoards()

	var winner, complete, cat int
	for _, b := range bs {
		if b.winner() {
			winner++

			if b.complete() {
				complete++
			}
		} else {
			cat++
		}
	}

	fmt.Printf("Board count: %d\n", len(bs))
	fmt.Printf("Winner count: %d [%d]\n", winner, complete)
	fmt.Printf("CAT count: %d\n", cat)

	/*
		for _, b := range bs {
			fmt.Printf("%s\n", b)
		}
	*/
}

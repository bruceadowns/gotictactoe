package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWinner(t *testing.T) {
	var tt = []struct {
		in     board
		expect bool
	}{
		{
			in:     board{X, Y, U, Y, X, U, U, Y, X},
			expect: true,
		},
		{
			in:     board{U, U, U, U, U, U, U, U, U},
			expect: false,
		},
	}

	for _, test := range tt {
		res := test.in.winner()
		if res != test.expect {
			t.Errorf("%v returned: %v expected: %v", test.in, res, test.expect)
		}
	}
}

func TestComplete(t *testing.T) {
	var tt = []struct {
		in     board
		expect bool
	}{
		{
			in:     board{U, U, U, U, U, U, U, U, U},
			expect: false,
		},
		{
			in:     board{X, Y, X, Y, X, Y, X, Y, X},
			expect: true,
		},
	}

	for _, test := range tt {
		res := test.in.complete()
		if res != test.expect {
			t.Errorf("%v returned: %v expected: %v", test.in, res, test.expect)
		}
	}
}

func TestPrint(t *testing.T) {
	var tt = []struct {
		in     board
		expect string
	}{
		{
			in:     board{U, U, U, U, U, U, U, U, U},
			expect: " | | \n-----\n | | \n-----\n | | \n",
		},
		{
			in:     board{X, Y, X, Y, X, Y, X, Y, X},
			expect: "x|y|x\n-----\ny|x|y\n-----\nx|y|x\n",
		},
	}

	for _, test := range tt {
		res := &bytes.Buffer{}
		fmt.Fprintf(res, "%s", test.in)
		t.Log(res)
		if res.String() != test.expect {
			t.Errorf("%v returned: %v expected: %v", test.in, res, test.expect)
		}
	}
}

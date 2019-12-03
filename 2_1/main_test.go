package main

import "testing"

func TestProcess(t *testing.T) {
	cases := []struct {
		in []int
		out []int
	}{
		{
			[]int{1,0,0,0,99},
			[]int{2,0,0,0,99},
		},
		{
			[]int{2,3,0,3,99},
			[]int{2,3,0,6,99},
		},
		{
			[]int{2,4,4,5,99,0},
			[]int{2,4,4,5,99,9801},
		},
		{
			[]int{1,1,1,4,99,5,6,0,99},
			[]int{30,1,1,4,2,5,6,0,99},
		},
	}

	for _, c := range cases {
		_ = process(c.in)

		for i := range c.out {
			if c.out[i] != c.in[i] {
				t.Errorf("Incorrect value at position %d\nExpected %d\nActual %d\n", i, c.out, c.in)
			}
		}
	}
}

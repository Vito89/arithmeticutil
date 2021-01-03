package arithmeticutil

import "testing"

func TestMin(t *testing.T) {
	cases := []struct {
		inA, inB, want int
	}{
		{0, 127, 0},
		{-1, 0, -1},
		{255, 0, 0},
	}
	for _, c := range cases {
		got := Min(c.inA, c.inB)
		if got != c.want {
			t.Errorf("Error with: %v, %v. Wanted: %v but got: %v", c.inA, c.inB, c.want, got)
		}
	}
}

package puck_test

import (
	"fmt"
	"testing"

	"github.com/DeedleFake/puck/puck"
)

func TestVercmp(t *testing.T) {
	lt := func(out int) bool {
		return out < 0
	}
	//gt := func(out int) bool {
	//	return out > 0
	//}
	eq := func(out int) bool {
		return out == 0
	}

	tests := []struct {
		v1, v2 puck.Version
		cmp    func(int) bool
	}{
		{
			puck.Version{1, []int{1, 0, 0}, 0},
			puck.Version{2, []int{1, 0, 0}, 0},
			lt,
		},
		{
			puck.Version{0, []int{1}, 0},
			puck.Version{0, []int{1, 0}, 0},
			lt,
		},
		{
			puck.Version{0, []int{1, 0}, 0},
			puck.Version{0, []int{1, 1}, 0},
			lt,
		},
		{
			puck.Version{0, []int{1, 1}, 0},
			puck.Version{0, []int{1, 1, 1}, 0},
			lt,
		},
		{
			puck.Version{0, []int{1, 1, 1}, 0},
			puck.Version{0, []int{1, 2}, 0},
			lt,
		},
		{
			puck.Version{0, []int{1, 2}, 0},
			puck.Version{0, []int{2, 0}, 0},
			lt,
		},
		{
			puck.Version{0, []int{2, 0}, 0},
			puck.Version{0, []int{3, 0, 0}, 0},
			lt,
		},
		{
			puck.Version{0, []int{1, 0, 0}, 0},
			puck.Version{0, []int{1, 0, 0}, 1},
			lt,
		},
		{
			puck.Version{1, []int{1, 0, 0}, 5},
			puck.Version{1, []int{1, 0, 0}, 5},
			eq,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("cmp(%v, %v)", test.v1, test.v2), func(t *testing.T) {
			out := puck.Vercmp(test.v1, test.v2)
			if !test.cmp(out) {
				t.Fatalf("Unexpected ouput: %v", out)
			}
		})
	}
}

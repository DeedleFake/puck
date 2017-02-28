package puck_test

import (
	"fmt"
	"testing"

	"github.com/DeedleFake/puck/puck"
)

func TestVercmp(t *testing.T) {
	tests := []struct {
		v1, v2 puck.Version
		out    int
	}{
		{
			puck.Version{1, "1.0.0", 0},
			puck.Version{2, "1.0.0", 0},
			puck.EpochLess,
		},
		{
			puck.Version{0, "1.0a", 0},
			puck.Version{0, "1.0b", 0},
			puck.VerLess,
		},
		{
			puck.Version{0, "1.0b", 0},
			puck.Version{0, "1.0beta", 0},
			puck.VerLess,
		},
		{
			puck.Version{0, "1.0beta", 0},
			puck.Version{0, "1.0p", 0},
			puck.VerLess,
		},
		{
			puck.Version{0, "1.0p", 0},
			puck.Version{0, "1.0pre", 0},
			puck.VerLess,
		},
		{
			puck.Version{0, "1.0pre", 0},
			puck.Version{0, "1.0rc", 0},
			puck.VerLess,
		},
		{
			puck.Version{0, "1.0rc", 0},
			puck.Version{0, "1.0", 0},
			puck.VerLess,
		},
		{
			puck.Version{0, "1.0", 0},
			puck.Version{0, "1.0.a", 0},
			puck.VerLess,
		},
		{
			puck.Version{0, "1.0.a", 0},
			puck.Version{0, "1.0.1", 0},
			puck.VerLess,
		},
		{
			puck.Version{0, "1", 0},
			puck.Version{0, "1.0", 0},
			puck.VerLess,
		},
		{
			puck.Version{0, "1.0", 0},
			puck.Version{0, "1.1", 0},
			puck.VerLess,
		},
		{
			puck.Version{0, "1.1", 0},
			puck.Version{0, "1.1.1", 0},
			puck.VerLess,
		},
		{
			puck.Version{0, "1.1.1", 0},
			puck.Version{0, "1.2", 0},
			puck.VerLess,
		},
		{
			puck.Version{0, "1.2", 0},
			puck.Version{0, "2.0", 0},
			puck.VerLess,
		},
		{
			puck.Version{0, "2.0", 0},
			puck.Version{0, "3.0.0", 0},
			puck.VerLess,
		},
		{
			puck.Version{0, "1.0.0", 0},
			puck.Version{0, "1.0.0", 1},
			puck.RelLess,
		},
		{
			puck.Version{1, "1.0.0", 5},
			puck.Version{1, "1.0.0", 5},
			puck.Equal,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("cmp(%v, %v)", test.v1, test.v2), func(t *testing.T) {
			out := puck.Vercmp(test.v1, test.v2)
			if out != test.out {
				t.Fatalf("Expected %v\nGot %v", test.out, out)
			}
		})
	}
}

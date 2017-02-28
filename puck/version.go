package puck

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// A Version holds information about a package's version. It is based
// on Arch Linux's pacman's versioning model. Unlike pacman, however,
// package release numbers start at 0, not 1.
type Version struct {
	Epoch uint
	Ver   string
	Rel   uint
}

func (ver Version) String() string {
	var buf bytes.Buffer

	if ver.Epoch != 0 {
		fmt.Fprintf(&buf, "%v:", ver.Epoch)
	}
	fmt.Fprintf(&buf, "%v-%v", ver.Ver, ver.Rel)

	return buf.String()
}

// These constants designate which part of a version a version
// comparison was decided by and what the outcome was. They are
// centered around 0 so that they can be ignored if their not
// necessary.
const (
	EpochLess = iota - 3
	VerLess
	RelLess
	Equal
	RelGreater
	VerGreater
	EpochGreater
)

func pcmp(p1, p2 string) int {
	n1, err := strconv.ParseUint(p1, 10, 0)
	an1 := err == nil

	n2, err := strconv.ParseUint(p2, 10, 0)
	an2 := err == nil

	switch {
	case an1 && !an2:
		if (n2 == 0) && (p1[0] != '0') {
			return VerLess
		}

		return VerGreater

	case an2 && !an1:
		if (n1 == 0) && (p2[0] != '0') {
			return VerLess
		}

		return VerGreater

	default:
		switch {
		case p1 < p2:
			return VerLess
		case p1 > p2:
			return VerGreater
		}
	}

	return Equal
}

func vercmp(v1, v2 string) int {
	p1 := strings.Split(v1, ".")
	p2 := strings.Split(v2, ".")

	for i := len(p1); i < len(p2); i++ {
		p1 = append(p1, "0")
	}
	for i := len(p2); i < len(p1); i++ {
		p2 = append(p2, "0")
	}

	for i := len(p1) - 1; i >= 0; i-- {
		if pc := pcmp(p1[i], p2[i]); pc != 0 {
			return pc
		}
	}

	return Equal
}

// Vercmp compares two versions. It follows the same rules as Arch
// Linux's vercmp utility. For more info, see
// https://www.archlinux.org/pacman/vercmp.8.html
func Vercmp(v1, v2 Version) int {
	switch {
	case v1.Epoch < v2.Epoch:
		return EpochLess
	case v1.Epoch > v2.Epoch:
		return EpochGreater
	}

	if vc := vercmp(v1.Ver, v2.Ver); vc != 0 {
		return vc
	}

	switch {
	case v1.Rel < v2.Rel:
		return RelLess
	case v1.Rel > v2.Rel:
		return RelGreater
	}

	return Equal
}

package puck

import (
	"bytes"
	"strconv"
)

// A Version holds information about a package's version. It is
// loosely based on Arch Linux's pacman's versioning model. There are
// a few significant differences, however:
//
// * The package version section of the version is based on semver,
//   rather than Arch's complicated alphanumeric versioning scheme.
//   Unlike semver, however, it allows as many sections to the version
//   as are necessary. This is to allow versions to be devised that
//   coorespond more appropriately with upstream version numbers.
// * Package release numbers start at 0, not 1.
type Version struct {
	Epoch int
	Ver   []int
	Rel   int
}

func (ver Version) String() string {
	var buf bytes.Buffer

	if ver.Epoch != 0 {
		buf.WriteString(strconv.FormatInt(int64(ver.Epoch), 10))
	}

	buf.WriteRune('v')
	buf.WriteString(strconv.FormatInt(int64(ver.Ver[0]), 10))
	for _, v := range ver.Ver[1:] {
		buf.WriteRune('.')
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	}

	buf.WriteRune('-')
	buf.WriteString(strconv.FormatInt(int64(ver.Rel), 10))

	return buf.String()
}

func vercmp(v1, v2 []int) int {
	for i := len(v1); i < len(v2); i++ {
		v1 = append(v1, -1)
	}
	for i := len(v2); i < len(v1); i++ {
		v2 = append(v2, -1)
	}

	for i := 0; i < len(v1); i++ {
		if vc := v1[i] - v2[i]; vc != 0 {
			return vc
		}
	}

	return 0
}

// Vercmp compares two versions. If the epochs are the same and the
// semver versions have different lengths, but are equal up until the
// end of the shorter one, the longer one is considered greater.
func Vercmp(v1, v2 Version) int {
	if ec := v1.Epoch - v2.Epoch; ec != 0 {
		return ec
	}

	if vc := vercmp(v1.Ver, v2.Ver); vc != 0 {
		return vc
	}

	return v1.Rel - v2.Rel
}

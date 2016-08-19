// Packager sorter provides convenience functions for sorting non-named,
// non-package level, and other types which do not implement sort.Interface.
//
// Inspiration comes from GitHub user Merovius. (See:
// https://github.com/golang/go/issues/16721#issuecomment-240384599)
package sorter

import "sort"

// LessFunc reports whether the element with index i should sort before the
// element with index j.
type LessFunc func(i, j int) bool

// SwapFunc swaps the elements with indexes i and j.
type SwapFunc func(i, j int)

// New returns a sort.Interface to allow sorting of non-named and non-package
// level types.
func New(len int, less LessFunc, swap SwapFunc) sort.Interface {
	return Sorter{N: len, L: less, S: swap}
}

// Sorter implements sort.Interface but allows sorting of non-named and
// non-package level types.
type Sorter struct {
	// N is the number of elements in the collection.
	N int
	// L reports whether the element with
	// index i should sort before the element with index j.
	L LessFunc
	// S swaps the elements with indexes i and j.
	S SwapFunc
}

// Len helps implement sort.Interface.
func (s Sorter) Len() int { return s.N }

// Less helps implement sort.Interface.
func (s Sorter) Less(i, j int) bool { return s.L(i, j) }

// Swap helps implement sort.Interface.
func (s Sorter) Swap(i, j int) { s.S(i, j) }

var _ sort.Interface = Sorter{}

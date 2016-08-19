package sorter

import (
	"math/rand"
	"sort"
	"testing"
)

var a [500]int

func init() {
	for i := 0; i < len(a); i++ {
		a[i] = rand.Int()
	}
}

func TestSorter(t *testing.T) {
	sort.Sort(Sorter{
		N: len(a),
		L: func(i, j int) bool { return a[i] < a[j] },
		S: func(i, j int) { a[i], a[j] = a[j], a[i] },
	})
	if !isSorted(a[:]) {
		t.Fatalf("not sorted: %#v", a)
	}
}

func isSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}

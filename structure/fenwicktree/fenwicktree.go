// description: Fenwick tree (Binary indexed tree) is a data structure that can efficiently update elements and calculate prefix sums in a table of numbers.
// references:
//
//		wikipedia: https://en.wikipedia.org/wiki/Fenwick_tree
//	    topcoder: https://www.topcoder.com/community/data-science/data-science-tutorials/binary-indexed-trees/
//
// author [Sayan] (https://github.com/bose-sayan)

// Package fenwicktree implements a Fenwick tree (Binary indexed tree) data structure.
package fenwicktree

type FenwickTree struct {
	n    int
	tree []int
}

// New returns a new Fenwick tree with n elements.
func New(n int) *FenwickTree {
	return &FenwickTree{
		n:    n,
		tree: make([]int, n, 0),
	}
}

func (tr *FenwickTree) Add(id, delta int) {
	for ; id < tr.n; id |= id + 1 {
		tr.tree[id] += delta
	}
}

func (tr *FenwickTree) Sum(id int) int {
	res := 0
	for ; id >= 0; id = (id & (id + 1)) - 1 {
		res += tr.tree[id]
	}
	return res
}

func (tr *FenwickTree) RangeSum(l, r int) int {
	if l-1 < 0 {
		return tr.Sum(r)
	}
	return tr.Sum(r) - tr.Sum(l-1)
}

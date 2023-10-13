package fenwicktree_test

import (
	"testing"

	"github.com/bose-sayan/Go/structure/fenwicktree"
)

func TestFenwickTree(t *testing.T) {
	// Initialize a new Fenwick tree with 10 elements
	tree := fenwicktree.New(10)

	// Test the Sum method
	if tree.Sum(5) != 0 {
		t.Errorf("Sum(5) should be 0")
	}

	// Test the Add method
	tree.Add(5, 10)
	if tree.Sum(5) != 10 {
		t.Errorf("Sum(5) should be 10")
	}

	// Test the Set method
	tree.Set(5, 5)
	if tree.Sum(5) != 5 {
		t.Errorf("Sum(5) should be 5")
	}

	// Test the RangeSum method
	if tree.RangeSum(2, 5) != 0 {
		t.Errorf("RangeSum(2, 5) should be 0")
	}

	// Test the AddRange method
	tree.AddRange(2, 5, 10)
	if tree.RangeSum(2, 5) != 10 {
		t.Errorf("RangeSum(2, 5) should be 10")
	}

	// Test the SetRange method
	tree.SetRange(2, 5, 5)
	if tree.RangeSum(2, 5) != 20 {
		t.Errorf("RangeSum(2, 5) should be 20")
	}
}

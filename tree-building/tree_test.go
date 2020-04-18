package tree

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

// Define a function Build(records []Record) (*Node, error)
// where Record is a struct containing int fields ID and Parent
// and Node is a struct containing int field ID and []*Node field Children.

var successTestCases = []struct {
	name     string
	input    []Record
	expected *Node
}{
	{
		name:     "empty input",
		input:    []Record{},
		expected: nil,
	},
	{
		name: "one node",
		input: []Record{
			{ID: 0},
		},
		expected: &Node{
			ID: 0,
		},
	},
	{
		name: "three nodes in order",
		input: []Record{
			{ID: 0},
			{ID: 1, Parent: 0},
			{ID: 2, Parent: 0},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				{ID: 1},
				{ID: 2},
			},
		},
	},
	{
		name: "three nodes in reverse order",
		input: []Record{
			{ID: 2, Parent: 0},
			{ID: 1, Parent: 0},
			{ID: 0},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				{ID: 1},
				{ID: 2},
			},
		},
	},
	{
		name: "more than two children",
		input: []Record{
			{ID: 3, Parent: 0},
			{ID: 2, Parent: 0},
			{ID: 1, Parent: 0},
			{ID: 0},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
		},
	},
	{
		name: "binary tree",
		input: []Record{
			{ID: 5, Parent: 1},
			{ID: 3, Parent: 2},
			{ID: 2, Parent: 0},
			{ID: 4, Parent: 1},
			{ID: 1, Parent: 0},
			{ID: 0},
			{ID: 6, Parent: 2},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				{
					ID: 1,
					Children: []*Node{
						{ID: 4},
						{ID: 5},
					},
				},
				{
					ID: 2,
					Children: []*Node{
						{ID: 3},
						{ID: 6},
					},
				},
			},
		},
	},
	{
		name: "unbalanced tree",
		input: []Record{
			{ID: 5, Parent: 2},
			{ID: 3, Parent: 2},
			{ID: 2, Parent: 0},
			{ID: 4, Parent: 1},
			{ID: 1, Parent: 0},
			{ID: 0},
			{ID: 6, Parent: 2},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				{
					ID: 1,
					Children: []*Node{
						{ID: 4},
					},
				},
				{
					ID: 2,
					Children: []*Node{
						{ID: 3},
						{ID: 5},
						{ID: 6},
					},
				},
			},
		},
	},
}

// Rules:
// 1. The ID number is always between 0 (inclusive) and the length of the record list (exclusive).
//    - lowest record must always be 0
//    - all ID's in range must be used, no gaps
//    - each record id must be unique
// 2. All records have a parent ID lower than their own ID, except for the root record, which has a parent ID that's equal to its own ID.
//    - record zero's parent must also be zero
//    - all other records must have a parent id lower than their own id

var failureTestCases = []struct {
	name  string
	input []Record
	err   error
}{
	{
		name: "lowest record must always be zero",
		input: []Record{
			{ID: 1, Parent: 0},
		},
		err: ErrorBadRootRecord,
	},
	{
		name: "no gaps in record ID's",
		input: []Record{
			{ID: 0},
			{ID: 2, Parent: 0},
		},
		err: ErrorNonContinuous,
	},
	{
		name: "each record id must be unique",
		input: []Record{
			{ID: 0},
			{ID: 1, Parent: 0},
			{ID: 2, Parent: 0},
			{ID: 2, Parent: 1},
		},
		err: ErrorNonContinuous,
	},
	{
		name: "root's parent must be zero",
		input: []Record{
			{ID: 0, Parent: -1},
		},
		err: ErrorBadRootRecord,
	},
	{
		name: "non zero records must have lower ID parent",
		input: []Record{
			{ID: 0},
			{ID: 1, Parent: 2},
			{ID: 2, Parent: 0},
		},
		err: ErrorBadParentID,
	},
}

func (n Node) String() string {
	return fmt.Sprintf("%d:%s", n.ID, n.Children)
}

func TestMakeTreeSuccess(t *testing.T) {
	for _, tt := range successTestCases {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Build(tt.input)
			if err != nil {
				var _ error = err
				t.Fatalf("Build for test case %q returned error %q. Error not expected.",
					tt.name, err)
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Fatalf("Build for test case %q returned %s but was expected to return %s.",
					tt.name, actual, tt.expected)
			}
		})
	}
}

func TestMakeTreeFailure(t *testing.T) {
	for _, tt := range failureTestCases {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Build(tt.input); err != tt.err {
				t.Fatalf("Expected error: %v\n Got error: %v", tt.err, err)
			}
		})
	}
}

func shuffleRecords(records []Record) []Record {
	gen := rand.New(rand.NewSource(42))
	newRecords := make([]Record, len(records))
	for i, idx := range gen.Perm(len(records)) {
		newRecords[i] = records[idx]
	}
	return newRecords
}

// Binary tree
func makeTwoTreeRecords() []Record {
	records := make([]Record, 1<<16)
	for i := range records {
		if i == 0 {
			records[i] = Record{ID: 0}
		} else {
			records[i] = Record{ID: i, Parent: i >> 1}
		}
	}
	return shuffleRecords(records)
}

var twoTreeRecords = makeTwoTreeRecords()

func BenchmarkTwoTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Build(twoTreeRecords)
	}
}

// Each node but the root node and leaf nodes has ten children.
func makeTenTreeRecords() []Record {
	records := make([]Record, 10000)
	for i := range records {
		if i == 0 {
			records[i] = Record{ID: 0}
		} else {
			records[i] = Record{ID: i, Parent: i / 10}
		}
	}
	return shuffleRecords(records)
}

var tenTreeRecords = makeTenTreeRecords()

func BenchmarkTenTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Build(tenTreeRecords)
	}
}

func makeShallowRecords() []Record {
	records := make([]Record, 10000)
	for i := range records {
		records[i] = Record{ID: i, Parent: 0}
	}
	return shuffleRecords(records)
}

var shallowRecords = makeShallowRecords()

func BenchmarkShallowTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Build(shallowRecords)
	}
}

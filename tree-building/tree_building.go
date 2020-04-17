package tree

import (
	"errors"
	"sort"
)

// Record is a single record in the tree with a reference to its parent
type Record struct {
	ID     int
	Parent int
}

// Node is a node within the tree which can contain nodes within it
type Node struct {
	ID       int
	Children []*Node
}

var (
	// ErrorRootHasParent root parent can not be >= 1
	ErrorRootHasParent = errors.New("root node must not have a parent")

	// ErrorNonContinuous root node must be zero with no duplicates and no gaps in the rest
	ErrorNonContinuous = errors.New("Nodes must be continuous")

	// ErrorParentHasHigherID parentID must not be greater than child id
	ErrorParentHasHigherID = errors.New("node can not be child of a higher id parent")

	// ErrorSelfParenting other than root childID must not equal parentID
	ErrorSelfParenting = errors.New("record can not have itself as parent")
)

// Build the tree
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make(map[int]*Node, len(records))
	for i, r := range records {
		if err := recordCheck(r, i); err != nil {
			return nil, err
		}

		node := &Node{ID: r.ID}
		nodes[r.ID] = node

		if r.ID != 0 {
			parentNode := nodes[r.Parent]
			parentNode.Children = append(parentNode.Children, node)
		}
	}

	return nodes[0], nil
}

func recordCheck(r Record, i int) error {
	if r.ID == 0 && r.Parent != 0 {
		return ErrorRootHasParent
	}

	if r.ID > 0 && r.ID == r.Parent {
		return ErrorSelfParenting
	}

	if r.Parent > r.ID {
		return ErrorParentHasHigherID
	}

	if r.ID != i {
		return ErrorNonContinuous
	}

	return nil
}

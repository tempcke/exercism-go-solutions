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
	// ErrorBadRootRecord root record must have id:0, parent:0
	ErrorBadRootRecord = errors.New("Lowest record ID must be zero with parent as zero")

	// ErrorNonContinuous root node must be zero with no duplicates and no gaps in the rest
	ErrorNonContinuous = errors.New("Nodes must be continuous")

	// ErrorBadParentID a node's parent id must be less than its own id
	ErrorBadParentID = errors.New("Except for record zero each record must have a lower id parent")
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
		if i == 0 && (r.ID != 0 || r.Parent != 0) {
			return nil, ErrorBadRootRecord
		}
		if r.ID != i {
			return nil, ErrorNonContinuous
		}
		if r.ID != 0 && r.Parent >= r.ID {
			return nil, ErrorBadParentID
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

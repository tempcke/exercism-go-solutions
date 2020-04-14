package tree

import (
	"fmt"
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

// Build the tree
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	// build node map
	nodes := make(map[int]*Node, len(records))
	for i, r := range records {
		if err := recordCheck(r, i); err != nil {
			return nil, err
		}
		nodes[r.ID] = &Node{ID: r.ID}

		if r.ID == 0 {
			continue
		}

		// append child
		// works only because the records are sorted
		// and rule: node can not be child of a higher id parent
		p := nodes[r.Parent]
		p.Children = append(p.Children, nodes[r.ID])
	}

	return nodes[0], nil
}

func recordCheck(r Record, i int) error {
	if r.ID > 0 && r.ID == r.Parent {
		return fmt.Errorf("record can not have itself as parent: %v", r)
	}

	if r.Parent > r.ID {
		return fmt.Errorf("node can not be child of a higher id parent: %v", r)
	}

	if r.ID != i {
		return fmt.Errorf("non-continuous: %v", r.ID)
	}

	return nil
}

package main

type Comparable interface {
	Less(than Comparable) bool
}

type Node struct {
	Data Comparable
}

func NewNode(data Comparable) *Node {
	return &Node{Data: data} }

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) insert(data Comparable) {
	if t.root == nil || t.root.Data.Less(data) {
		t.root = NewNode(data)
	}
}

type Row struct {
	Row []string
}

func NewRow(row[] string) *Row {
	return &Row{Row: row}
}

func (r Row) Less(other Comparable) bool {
	return r.Row[0] < other.(Row).Row[0]
}

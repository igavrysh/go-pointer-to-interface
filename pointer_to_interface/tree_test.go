package main

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	d := []string{"42"}
	tree := NewTree()
	tree.insert(*NewRow(d))

	if !reflect.DeepEqual(tree.root.Data.(Row).Row, d) {
		t.Error("returned elements are not matching")
	}
}

func TestInsert2x(t *testing.T) {
	d1 := []string{"42"}
	d2 := []string{"99"}
	tree := NewTree()
	tree.insert(*NewRow(d1))
	tree.insert(*NewRow(d2))

	if !reflect.DeepEqual(tree.root.Data.(Row).Row, d2) {
		t.Error("returned elements are not matching")
	}
}

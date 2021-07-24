package main

import "fmt"

func main() {
	t := NewTree()
	t.insert(NewRow([]string{"123"}))
	fmt.Printf("%v\n", t.root.Data.(*Row).Row)
}

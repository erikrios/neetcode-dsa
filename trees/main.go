package main

import (
	"fmt"
	"trees/tree/binary"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	bst := binary.New()
	spew.Dump(bst)
	fmt.Println("-----")
	bst.Insert(10)
	spew.Dump(bst)
	fmt.Println("-----")
	bst.Insert(5)
	spew.Dump(bst)
	fmt.Println("-----")
	bst.Insert(12)
	spew.Dump(bst)
	fmt.Println("-----")
	bst.Insert(8)
	spew.Dump(bst)
	fmt.Println(bst.Inorder())
	fmt.Println(bst.Preorder())
	fmt.Println(bst.Postorder())
}

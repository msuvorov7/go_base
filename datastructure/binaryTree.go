package main

import (
	"fmt"
)

type Node struct {
	val int
	left *Node
	right *Node
}


func main() {
	tree := Node{5, nil, nil}
	tree.left = &Node{3, nil, nil}
	tree.right = &Node{7, nil, nil}
	fmt.Println(tree.val)
	fmt.Println(tree.left.val, tree.right.val)

}
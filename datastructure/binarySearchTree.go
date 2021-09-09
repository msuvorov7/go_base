package main

import (
	"fmt"
	"errors"
	"strconv"
)

type Node struct {
	data int
	left *Node
	right *Node
}

func (n *Node) insert(val int) error {
	if n == nil {
		return errors.New("tree is nil")
	}
	if val == n.data {
		return errors.New("value is already exists")
	}
	if (val < n.data) && (n.left == nil) {
		n.left = &Node{val, nil, nil}
		return nil
	} else if (val < n.data) {
		return n.left.insert(val)
	}
	if (val > n.data) && (n.right == nil) {
		n.right = &Node{val, nil, nil}
		return nil
	} else if (val > n.data) {
		return n.right.insert(val)
	}
	return nil
}

func (n *Node) inorder() {
	// вывод дерева в отсортированном порядке
	if n != nil {
		n.left.inorder()
		fmt.Print(n.data)
		n.right.inorder()
	}
}

func (n *Node) search(val int) *Node {
	if (n == nil) || (n.data == val) {
		return n
	}
	if (val < n.data) {
		return n.left.search(val)
	} else {
		return n.right.search(val)
	}
}

func (n *Node) toString(res *[]string) {
	if n != nil {
		n.left.toString(res)
		*res = append(*res, "(")
		*res = append(*res, strconv.Itoa(n.data))
		n.right.toString(res)
		*res = append(*res, ")")
	}
}

// func (n *Node) getMin() *Node {
// 	curr := n
// 	if curr == nil {
// 		return nil
// 	}
// 	for curr.left != nil {
// 		curr = curr.left // тоже работает
// 	} 
// 	return curr
// }

func getMin(root *TreeNode) *TreeNode {
    curr := root
    if curr == nil {
        return curr
    }
    if curr.Left != nil {
        curr = getMin(curr.Left)
    }
    return curr
}

func (n *Node) getMax() *Node {
	curr := n
	if curr == nil {
		return nil
	}
	for curr.right != nil {
		curr = curr.right.getMin()
	} 
	return curr
}

func (root *Node) removeNode(key int) *Node {
	if root == nil {
		return root
	}
	if key < root.data {
		root.left = root.left.removeNode(key)
	} else if key > root.data {
		root.right = root.right.removeNode(key)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		temp := root.right.getMin()
		root.data = temp.data
		root.right = root.right.removeNode(temp.data)
	}
	return root
}


func main() {
	tree := &Node{5, nil, nil}
	tree.insert(2)
	tree.insert(4)
	tree.insert(7)
	tree.insert(1)
	tree.insert(3)
	tree.insert(9)
	tree.inorder()
	fmt.Println()

	finded := tree.search(4)
	fmt.Println(finded.data == 4)
	str := make([]string, 0)
	tree.toString(&str)
	fmt.Println(str)
	fmt.Println("min ", tree.getMin().data)
	fmt.Println("max ", tree.getMax().data)

	tree = tree.removeNode(7)
	tree.inorder()
	fmt.Println()
}
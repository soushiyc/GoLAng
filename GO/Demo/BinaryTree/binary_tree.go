package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value int
	left *node
	right *node
}


/* Binary Search Tree


						3

				2				5

			1				4		6


*/

type bst struct {
	root *node
	len int

}

func (n node) String() string{
	return strconv.Itoa(n.value)
}

func (b bst) String() string{
	sb := strings.Builder{}
	b.inorderTraversal(&sb)
	return sb.String()
}

func (b bst) inorderTraversal(sb *strings.Builder){
	b.inorderTraversalByNode(sb, b.root)
}

func (b bst) inorderTraversalByNode(sb *strings.Builder, root *node){
	if root == nil{
		return
	}
	b.inorderTraversalByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s", root))
	b.inorderTraversalByNode(sb, root.right)
}

func (b *bst) add(value int) {
	b.root = b.addByNode(b.root, value)
	b.len++
}

func (b *bst) addByNode(root *node, value int) *node {
	if root == nil{
		return &node{value: value}
	}

	if value < root.value{
		root.left = b.addByNode(root.left, value)
	} else {
		root.right = b.addByNode(root.right, value)
	}
	return root
}

func (b bst) search(value int) (*node, bool) {
	return b.searchByNode(b.root, value)
}

func (b bst) searchByNode(root *node, value int) (*node, bool) {
	if root == nil{
		return nil, false
	}
	if value == root.value{
		return root, true
	} else if value < root.value{
		return b.searchByNode(root.left, value)
	} else {
		return b.searchByNode(root.right, value)
	}
}

/*
Two methods for remove operations:
1: Replacing the searched value with the largest value in the right subtree. And Removing the largest value in the right subtree.
2: Replacing the searched value with the smallest value in the left subtree. And Removing the smallest value in the left subtree.
In this code, We'll be performing the second method
*/

func (b *bst) remove(value int) {
	b.removeByNode(b.root, value)
}

func (b *bst) removeByNode(root *node, value int) *node {
	if root == nil{
	return nil
	}
	if value < root.value{
		root.left = b.removeByNode(root.left, value)
	} else if value > root.value{
		root.right = b.removeByNode(root.right, value)
	} else{
		if root.left == nil{
			return root.right
		} else {
			temp := root.left
			for temp.right != nil{
				temp = temp.right
			}
			root.value = temp.value
			root.left = b.removeByNode(root.left, value)
		}
	}
	return root
}

func main()	{
	n := 	&node{5, nil, nil}
	n.left = &node{2, nil, nil}
	n.right = &node{6, nil, nil}

	b := bst{
		root: n,
		len: 3,
	}
	fmt.Println(b)
	fmt.Println()
	b.add(7)
	b.add(1)
	b.add(3)
	fmt.Println(b)
	b.add(4)
	fmt.Println(b)
	fmt.Println()

	fmt.Println(b.search(1))
	fmt.Println()

	b.remove(3)
	fmt.Println(b)
	b.remove(6)
	fmt.Println(b)
	b.remove(5)
	fmt.Println(b)

}
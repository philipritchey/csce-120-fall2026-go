package main

import (
	"fmt"
	"math/rand/v2"
)

type FullBinaryTreeNode struct {
	Value       int
	Left, Right *FullBinaryTreeNode
}

func main() {
	tree := MakeRandomFullBinaryTree(10)
	PrintTree(tree)
	DepthFirstSearchTree(tree, 42)
}

func DepthFirstSearchTree(root *FullBinaryTreeNode, target int) bool {
	if root == nil {
		// dead end
		return false
	}
	if root.Value == target {
		// success
		fmt.Printf("found %d!\n", target)
		return true
	}
	// try searching left
	if DepthFirstSearchTree(root.Left, target) {
		fmt.Printf("left @ %d\n", root.Value)
		return true
	}
	// then try searching right
	if DepthFirstSearchTree(root.Right, target) {
		fmt.Printf("right @ %d\n", root.Value)
		return true
	}
	// neither child was a success, failure as a parent
	return false
}

// make a full binary tree with n leaf nodes and (n-1) random-valued internal nodes
func MakeRandomFullBinaryTree(n int) *FullBinaryTreeNode {
	if n == 0 {
		return nil
	}
	nodes := make([]*FullBinaryTreeNode, n)
	for i := range n {
		nodes[i] = new(FullBinaryTreeNode{Value: rand.IntN(n * n)})
	}
	for len(nodes) > 1 {
		// pick two nodes at random and make a new FBT with a random-valued root
		i := rand.IntN(len(nodes))
		j := rand.IntN(len(nodes)-1) + 1
		nodes[0], nodes[i] = nodes[i], nodes[0]
		nodes[1], nodes[j] = nodes[j], nodes[1]
		node := new(FullBinaryTreeNode{
			Value: rand.IntN(n * n),
			Left:  nodes[0],
			Right: nodes[1],
		})
		nodes = nodes[2:]
		nodes = append(nodes, node)
	}
	return nodes[0]
}

// PrintTree prints a binary tree sideways with visible edges.
// Right children appear above the node, left children below.
func PrintTree(root *FullBinaryTreeNode) {
	printTree(root, "", true, true)
}

func printTree(node *FullBinaryTreeNode, prefix string, isTail, isRoot bool) {
	if node == nil {
		return
	}

	if node.Right != nil {
		newPrefix := prefix
		if isTail && !isRoot {
			newPrefix += " │   "
		} else {
			newPrefix += "     "
		}
		printTree(node.Right, newPrefix, false, false)
	}

	var corner string
	if isRoot {
		corner = "     "
	} else if isTail {
		corner = " └── "
	} else {
		corner = " ┌── "
	}
	fmt.Println(prefix + corner + fmt.Sprint(node.Value))

	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "     "
		} else {
			newPrefix += " │   "
		}
		printTree(node.Left, newPrefix, true, false)
	}
}

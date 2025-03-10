``` go
package main

import "fmt"

type Node struct {
    Value       int
    Left, Right *Node
}

// Tree represents the binary search tree.
type Tree struct {
	Root *Node
}

// Insert adds a new value to the tree.
func (t *Tree) Insert(value int) {
	t.Root = insertNode(t.Root, value)
}

// insertNode recursively finds the correct location for the new value.
func insertNode(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	}
	if value < root.Value {
		root.Left = insertNode(root.Left, value)
	} else if value > root.Value {
		root.Right = insertNode(root.Right, value)
	}
	return root
}

// Search looks for a value in the tree and returns the node if found.
func (t *Tree) Search(value int) *Node {
	return searchNode(t.Root, value)
}

// searchNode recursively searches for a node with the given value.
func searchNode(root *Node, value int) *Node {
	if root == nil || root.Value == value {
		return root
	}
	if value < root.Value {
		return searchNode(root.Left, value)
	}
	return searchNode(root.Right, value)
}

// Delete removes a value from the tree.
func (t *Tree) Delete(value int) {
	t.Root = deleteNode(t.Root, value)
}

// deleteNode removes a node with the given value and adjusts the tree accordingly.
func deleteNode(root *Node, value int) *Node {
	if root == nil {
		return nil
	}
	if value < root.Value {
		root.Left = deleteNode(root.Left, value)
	} else if value > root.Value {
		root.Right = deleteNode(root.Right, value)
	} else {
		// Node to be deleted found.
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		// Node with two children: get the inorder successor (smallest in the right subtree).
		minRight := findMin(root.Right)
		root.Value = minRight.Value
		root.Right = deleteNode(root.Right, minRight.Value)
	}
	return root
}

// findMin returns the node with the minimum value in the given subtree.
func findMin(root *Node) *Node {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// InOrder performs an in-order traversal of the tree.
func (t *Tree) InOrder() {
	inOrder(t.Root)
	fmt.Println()
}

// inOrder recursively traverses the tree in-order.
func inOrder(root *Node) {
	if root != nil {
		inOrder(root.Left)
		fmt.Printf("%d ", root.Value)
		inOrder(root.Right)
	}
}

// PreOrder performs a pre-order traversal of the tree.
func (t *Tree) PreOrder() {
	preOrder(t.Root)
	fmt.Println()
}

// preOrder recursively traverses the tree pre-order.
func preOrder(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.Value)
		preOrder(root.Left)
		preOrder(root.Right)
	}
}

// PostOrder performs a post-order traversal of the tree.
func (t *Tree) PostOrder() {
	postOrder(t.Root)
	fmt.Println()
}

// postOrder recursively traverses the tree post-order.
func postOrder(root *Node) {
	if root != nil {
		postOrder(root.Left)
		postOrder(root.Right)
		fmt.Printf("%d ", root.Value)
	}
}

// LevelOrder performs a level-order (breadth-first) traversal of the tree.
func (t *Tree) LevelOrder() {
	if t.Root == nil {
		return
	}
	queue := []*Node{t.Root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node.Value)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Println()
}
```

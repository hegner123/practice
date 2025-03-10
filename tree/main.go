package main

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(value int) {
	t.Root = insertNode(t.Root, value)
}

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

func (t *Tree) Search(value int) *Node {
	return searchNode(t.Root, value)
}

func searchNode(root *Node, value int) *Node {
	if root == nil || root.Value == value {
		return root
	}
	if value < root.Value {
		return searchNode(root.Left, value)
	}
	return searchNode(root.Right, value)

}

func (t *Tree) Delete(value int) {
	t.Root = deleteNode(t.Root, value)
}

func deleteNode(root *Node, value int) *Node {
	if root == nil {
		return nil
	}
	if value < root.Value {
		deleteNode(root.Left, value)
	} else if value > root.Value {
		deleteNode(root.Right, value)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		minRight := findMin(root.Right)
		root.Value = minRight.Value
		root.Right = deleteNode(root.Right, value)

	}
	return root
}

func findMin(root *Node) *Node {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func (t *Tree) InOrder() {
	inOrder(t.Root)
	fmt.Println()
}

func inOrder(root *Node) {
	if root != nil {
		inOrder(root.Left)
		fmt.Printf("%d ", root.Value)
		inOrder(root.Right)
	}
}

func (t *Tree) PreOrder() {
	preOrder(t.Root)
	fmt.Println()
}

func preOrder(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.Value)
		preOrder(root.Left)
		preOrder(root.Right)
	}

}
func (t *Tree) PostOrder() {
	postOrder(t.Root)
	fmt.Println()

}
func postOrder(root *Node) {
	if root != nil {
		postOrder(root.Left)
		postOrder(root.Right)
		fmt.Printf("%d ", root.Value)
	}

}
func (t *Tree) LevelOrder() {
	if t.Root == nil {
		return
	}
	queue := []*Node{t.Root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[:1]
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

func main() {

}

package main

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T){
    // Create a root node.
    root := &Tree{}
    // Insert some values.
    values := []int{5, 15, 3, 7, 12, 18}
    for _, v := range values {
        root.Insert(v)
    }

    // Print the tree using in-order traversal.
    fmt.Print("In-order traversal: ")
    root.InOrder()
    fmt.Print("Post-order traversal: ")
    root.PostOrder()
    fmt.Print("Pre-order traversal: ")
    root.PreOrder()
    fmt.Println()

    fmt.Println("Delete 5 ")
    root.Delete(5)
    fmt.Print("In-order traversal: ")
    root.InOrder()
    fmt.Print("Post-order traversal: ")
    root.PostOrder()
    fmt.Print("Pre-order traversal: ")
    root.PreOrder()
    fmt.Println()

    fmt.Println("Insert 70 ")
    root.Insert(70)
    fmt.Print("In-order traversal: ")
    root.InOrder()
    fmt.Print("Post-order traversal: ")
    root.PostOrder()
    fmt.Print("Pre-order traversal: ")
    root.PreOrder()
    fmt.Println()

    fmt.Printf("search value 7=%d\n",root.Search(7).Value)

}

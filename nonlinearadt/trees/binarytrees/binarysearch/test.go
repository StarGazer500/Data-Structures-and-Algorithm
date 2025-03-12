package binarysearch

import (

    "math/rand"
    "time"
    "fmt"
)


type Number int

func (num Number) String() string {
    return fmt.Sprintf("%d", num)
}
type Float float64
func (num Float) String() string {
    return fmt.Sprintf("%0.1f", num)
}
func inorderOperator(val Float) {
    fmt.Println(val.String())
}


func TestBinarySearchTree(){
	// rand.Seed(time.Now().UnixNano())
    // Generate a random search tree
    randomSearchTree := BinarySearchTree[Float]{nil, 0}
    for i := 0; i < 30; i++ {
        rn := 1.0 + 99.0 * rand.Float64()
        randomSearchTree.Insert(Float(rn))
    }
	time.Sleep(3 * time.Second)
	randomSearchTree.InOrderTraverse(inorderOperator)
    min := randomSearchTree.Min()
    max, _ := randomSearchTree.Max()
    fmt.Printf("\nMinimum value in random search tree is %0.1f  \nMaximum value in random search tree is %0.1f", *min, *max)
	start := time.Now()
    tree := BinarySearchTree[Number]{nil, 0}
    for val := 0; val < 100_000; val++ {
        tree.Insert(Number(val))
    }
    elapsed := time.Since(start)
    _, ht := tree.Max()

	fmt.Printf("\nTime to build BST tree with 100,000 nodes in sequential order: %s. Height of tree: %d", elapsed, ht)
}
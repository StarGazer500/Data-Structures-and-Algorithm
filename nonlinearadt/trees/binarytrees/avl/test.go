package avl

import (
    
    "fmt"
    "math/rand"
    "time"
	"sort"
)

func inorderOperator(val Float) {
    val *= val
    fmt.Println(val.String())
}

type Float float64
func (num Float) String() string {
    return fmt.Sprintf("%0.1f", num)
}
type Integer int
func (num Integer) String() string {
    return fmt.Sprintf("%d", num)
}

func TestAvlTree(){
	// rand.Seed(time.Now().UnixNano())
    // Generate a random search tree
    randomSearchTree := AVLTree[Float]{nil, 0}
    for i := 0; i < 30; i++ {
        rn := 1.0 + 99.0 * rand.Float64()
        randomSearchTree.Insert(Float(rn))
    }
    time.Sleep(3 * time.Second)
  
    randomSearchTree.InOrderTraverse(inorderOperator)
    min := randomSearchTree.Min()
    max := randomSearchTree.Max()
    fmt.Printf("\nMinimum value in tree is %0.1f  Maximum value in tree is %0.1f", *min, *max)

    start := time.Now()
    tree := AVLTree[Integer]{nil, 0}
    for val := 0; val < 100_0000; val++ {
        tree.Insert(Integer(val))
    }
    elapsed := time.Since(start)

	fmt.Printf("\nTime to build AVL tree with 100,0000 nodes: %s.  Height of tree: %d", elapsed, tree.Height())
	numbers := make([]int, 100_000)
	for i := 0; i < 100_000; i++ {
	numbers[i] = i
	}
	start = time.Now()
	sort.Ints(numbers)
	elapsed = time.Since(start)
	fmt.Printf("\nTime to sort 100_000 ints: %s", elapsed)
}
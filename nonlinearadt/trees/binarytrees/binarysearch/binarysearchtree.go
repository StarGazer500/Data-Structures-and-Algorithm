package binarysearch
type ordered interface {
    ~int | ~float64 | ~string
}
type BinarySearchTree[T OrderedStringer] struct {
    Root *Node[T]
    NumNodes int
}
type Node[T OrderedStringer] struct {
    Value T
    Left *Node[T]
    Right *Node[T]
}
type OrderedStringer interface {
    ordered
    String() string
}

func insertNode[T OrderedStringer](node, newNode *Node[T]) {
    if newNode.Value < node.Value {
        if node.Left == nil {
            node.Left = newNode
        } else {
            insertNode(node.Left, newNode)
        }
    } else {
        if node.Right == nil {
            node.Right = newNode
        } else {
            insertNode(node.Right, newNode)
        }
    }
}

func (bst *BinarySearchTree[T]) Insert(newValue T) {
    if bst.Search(newValue) == false { // newValue not in existing tree
        n := &Node[T]{newValue, nil, nil}
        if bst.Root == nil { // First value in bst
            bst.Root = &Node[T]{newValue, nil, nil}
        } else {
            insertNode(bst.Root, n)
        }
        bst.NumNodes += 1
    }

}

func search[T OrderedStringer](n *Node[T], value T) bool {
    if n == nil {
        return false
    }
    if value < n.Value {
        return search(n.Left, value)
    }
    if value > n.Value {
        return search(n.Right, value)
    }
    return true
}

func (bst *BinarySearchTree[T]) Search(value T) bool {
    return search(bst.Root, value)
}

func deleteNode[T OrderedStringer](node *Node[T], value T) *Node[T] {
    if node == nil {
        return nil
    }
    
    if value < node.Value {
        node.Left = deleteNode(node.Left, value)
        return node
    }
    if value > node.Value {
        node.Right = deleteNode(node.Right, value)
        return node
    }
    
    // Node with value to delete found
    
    // Case 1: No children
    if node.Left == nil && node.Right == nil {
        return nil
    }
    
    // Case 2: One child
    if node.Left == nil {
        return node.Right
    }
    if node.Right == nil {
        return node.Left
    }
    
    // Case 3: Two children
    // Find the minimum value in the right subtree (successor)
    successor := node.Right
    for successor.Left != nil {
        successor = successor.Left
    }
    
    // Copy successor value to current node
    node.Value = successor.Value
    
    // Delete the successor
    node.Right = deleteNode(node.Right, successor.Value)
    return node
}

func (bst *BinarySearchTree[T]) Delete(value T) {
    if bst.Search(value) == true {
        deleteNode(bst.Root, value)
        bst.NumNodes -= 1
    }
}

func inOrderTraverse[T OrderedStringer](n *Node[T], op func(T)) {
    if n != nil {
        inOrderTraverse(n.Left, op)
        op(n.Value)
        inOrderTraverse(n.Right, op)
    }
}
func (bst *BinarySearchTree[T]) InOrderTraverse(op func(T)) {
    inOrderTraverse(bst.Root, op)
}


func (bst *BinarySearchTree[T]) Min() *T {
    node := bst.Root
    if node == nil {
        return nil
    }
    for {
        if node.Left == nil {
            return &node.Value
        }
        node = node.Left
    }
}


func (bst *BinarySearchTree[T]) Max() (*T, int) { // second return value is
                                                  // height
    node := bst.Root
    height := 1
    if node == nil {
        return nil, 0
    }
    for {
        if node.Right == nil {
            return &node.Value, height
        }
        height += 1
        node = node.Right
    }
}

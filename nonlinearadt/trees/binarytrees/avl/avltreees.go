package avl


type ordered interface {
    ~int | ~float64 | ~string
}
type AVLTree[T OrderedStringer] struct {
    Root *Node[T]
    NumNodes int
}


type Node[T OrderedStringer] struct {
    Value T
    Left *Node[T]
    Right *Node[T]
    Ht int
}
type OrderedStringer interface {
    ordered
    String() string
}

// Methods

func (n *Node[T]) updateHeight() {
    max := func (a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    n.Ht = max(n.Left.Height(), n.Right.Height()) + 1
}


func rightRotate[T OrderedStringer](x *Node[T]) *Node[T] {
    y := x.Left
    t := y.Right
    y.Right = x
    x.Left = t
    x.updateHeight()
    y.updateHeight()
    return y
}

func leftRotate[T OrderedStringer](x *Node[T]) *Node[T] {
    y := x.Right
    t := y.Left
    y.Left = x
    x.Right = t
    x.updateHeight()
    y.updateHeight()
  return y
}

func rotateInsert[T OrderedStringer](node *Node[T], val T) *Node[T] {
    node.updateHeight()
    bFactor := node.balanceFactor()
    if bFactor > 1 && val < node.Left.Value {
      return rightRotate(node)
    }
    if bFactor < -1 && val > node.Right.Value {
      return leftRotate(node)
    }
    if bFactor > 1 && val > node.Left.Value {
      node.Left = leftRotate(node.Left)
      return rightRotate(node)
    }
    if bFactor < -1 && val < node.Right.Value {
      node.Right = rightRotate(node.Right)
      return leftRotate(node)
    }
    return node
  }

func insertNode[T OrderedStringer](node *Node[T], val T) *Node[T] {
    // if there's no node, create one
    if node == nil {
        return newNode(val)
    }
    // if value is greater than current node's value, insert to the right
    if val > node.Value {
        right := insertNode(node.Right, val)
        node.Right = right
    }
    // if value is less than current node's value, insert to the left
    if val < node.Value {
        left:= insertNode(node.Left, val)
        node.Left = left
    }
    return rotateInsert(node, val)
}

func (avl *AVLTree[T]) Insert(newValue T) {
    if avl.Search(newValue) == false { // newValue is not in existing tree
        avl.Root = insertNode(avl.Root, newValue)
        avl.NumNodes += 1
    }
}

func largest[T OrderedStringer](node *Node[T]) *Node[T] {
    if node == nil {
        return nil
    }
    if node.Right == nil {
        return node
    }
    return largest(node.Right)
}

func deleteNode[T OrderedStringer](node *Node[T], val T) *Node[T] {
    if node == nil {
        return nil
    }

    if val > node.Value {
        right := deleteNode(node.Right, val)
        node.Right = right
    } else if val < node.Value {
        left := deleteNode(node.Left, val)
        node.Left = left
    } else {
        if node.Left != nil && node.Right != nil {
            // has 2 children, find the successor
            successor := largest(node.Left)
            value := successor.Value
            // remove the successor
            left := deleteNode(node.Left, value)
            node.Left = left
            // copy the successor value to the current node
            node.Value = value
        } else if node.Left != nil || node.Right != nil {
            // has 1 child
            // move the child position to the current node
            if node.Left != nil {
                node = node.Left
            } else {
                node = node.Right
            }
        } else if node.Left == nil && node.Right == nil {
            // has no child
            // simply remove the node
            node = nil
        }
    }

    if node == nil {
        return nil
    }

    return rotateDelete(node)
}


func (avl *AVLTree[T]) Delete(value T) {
    if avl.Search(value) == true {
        avl.Root = deleteNode(avl.Root, value)
        avl.NumNodes -= 1
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

func (avl *AVLTree[T]) Search(value T) bool {
    return search(avl.Root, value)
}

func (n *Node[T]) Height() int {
    if n == nil {
        return 0
    } else {
        return n.Ht
    }
}

func (avl *AVLTree[T]) Height() int {
    return avl.Root.Height()
}

func inOrderTraverse[T OrderedStringer](n *Node[T], op func(T)) {
    if n != nil {
        inOrderTraverse(n.Left, op)
        op(n.Value)
        inOrderTraverse(n.Right, op)
    }
}

func (avl *AVLTree[T]) InOrderTraverse(f func(T)) {
    inOrderTraverse(avl.Root, f)
}

func (avl *AVLTree[T]) Min() *T {
    node := avl.Root
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


func (avl *AVLTree[T]) Max() *T {
    node := avl.Root
    if node == nil {
        return nil
    }
    for {
        if node.Right == nil {
            return &node.Value
        }
        node = node.Right
    }
}


func (n *Node[T]) balanceFactor() int {
    if n == nil {
        return 0
    }
    return n.Left.Height() - n.Right.Height()
}



func newNode[T OrderedStringer](val T) *Node[T] {
    return &Node[T] {
        Value:  val,
        Left:   nil,
        Right:  nil,
        Ht: 1,
    }
}

func rotateDelete[T OrderedStringer](node *Node[T]) *Node[T] {
    node.updateHeight()
    bFactor := node.balanceFactor()
    if bFactor > 1 && node.Left.balanceFactor() >= 0 {
        return rightRotate(node)
    }
    if bFactor > 1 && node.Left.balanceFactor() < 0 {
        node.Left = leftRotate(node.Left)
        return rightRotate(node)
    }
    if bFactor < -1 && node.Right.balanceFactor() <= 0 {
        return leftRotate(node)
    }
    if bFactor < -1 && node.Right.balanceFactor() > 0 {
        node.Right = rightRotate(node.Right)
        return leftRotate(node)
    }
    return node
}
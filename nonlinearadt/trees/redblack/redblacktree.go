package redblack

type ordered interface {
    ~int | ~float64 | ~string
}


type OrderedStringer interface {
    ordered
    String() string
}
type Node[T OrderedStringer] struct {
    value T
    red bool
    parent *Node[T]
    left *Node[T]
    right *Node[T]
}

type RedBlackTree[T OrderedStringer] struct {
    count int
    root *Node[T]
}
func NewTree[T OrderedStringer](value T) *RedBlackTree[T] {
    return &RedBlackTree[T]{1, &Node[T]{value, false, nil, nil, nil}}
}

func (tree *RedBlackTree[T]) Insert(value T) {
    if tree.root == nil { // Empty tree
        tree.root = &Node[T]{value, false, nil, nil, nil}
        tree.count += 1
        return
    }
    parent, nodeDirection := tree.findParent(value)
    if nodeDirection == "" {
        return
    }
    newNode := Node[T]{value, true, parent, nil, nil}
    if nodeDirection == "L" {
        parent.left = &newNode
		} else {
			parent.right = &newNode
		}
		tree.checkReconfigure(&newNode)
		tree.count += 1
	}

	func (tree *RedBlackTree[T]) IsPresent(value T, node*Node[T]) bool {
    if node == nil {
        return false
    }
    if value < node.value {
        return tree.IsPresent(value, node.left)
    }
    if value > node.value {
        return tree.IsPresent(value, node.right)
    }
    return true
}


func (tree *RedBlackTree[T]) findParent(value T)(*Node[T], string) {
    return search(value, tree.root)
}
func (tree *RedBlackTree[T]) checkReconfigure(node *Node[T]) {
    var nodeDirection, parentDirection, rotation string
    var uncle *Node[T]
    parent := node.parent
    value :=  node.value
    if parent == nil || parent.parent == nil ||
                node.red == false || parent.red == false {
        return
    }
    grandfather := parent.parent
	if value < parent.value {
        nodeDirection = "L"
    } else {
        nodeDirection = "R"
    }
    if grandfather.value > parent.value {
        parentDirection = "L"
    } else {
        parentDirection = "R"
    }
    if parentDirection == "L" {
        uncle = grandfather.right
    } else {
        uncle = grandfather.left
    }
    rotation = nodeDirection + parentDirection
    if uncle == nil || uncle.red == false {
        if rotation == "LL" {
            tree.rightRotate(node, parent, grandfather, true)
        } else if rotation == "RR" {
            tree.leftRotate(node, parent, grandfather, true)
        } else if rotation == "LR" {
            tree.rightRotate(nil, node, parent, false)
            tree.leftRotate(parent, node, grandfather, true)
            node, parent = parent, node
        } else if rotation == "RL" {
            tree.leftRotate(nil, node, parent, false)
            tree.rightRotate(parent, node, grandfather, true)
        }
    } else {
        tree.modifyColor(grandfather)
    }
}

func (tree *RedBlackTree[T]) leftRotate(node, parent, grandfather *Node[T],  modifyColor bool) {
	greatgrandfather := grandfather.parent
    tree.updateParent(parent, grandfather, greatgrandfather)
    oldLeft := parent.left
    parent.left = grandfather
    grandfather.parent = parent
    grandfather.right  = oldLeft
    if oldLeft != nil {
        oldLeft.parent = grandfather
    }
    if modifyColor == true {
        parent.red = false
        node.red = true
        grandfather.red = true
    }
}

func (tree *RedBlackTree[T]) rightRotate(node, parent,grandfather *Node[T],  modifyColor bool) {
greatgrandfather := grandfather.parent
tree.updateParent(parent, grandfather,
				  greatgrandfather)
oldRight := parent.right
parent.right = grandfather
grandfather.parent = parent
grandfather.left = oldRight
if oldRight != nil {
	oldRight.parent = grandfather
}
if modifyColor == true {
	parent.red = false
	node.red = true
	grandfather.red = true
}
}

func (tree *RedBlackTree[T]) modifyColor(grandfather *Node[T]) {
	grandfather.right.red = false
	grandfather.left.red = false
	if grandfather != tree.root {
		grandfather.red = true
	}
	tree.checkReconfigure(grandfather)
}


func (tree *RedBlackTree[T]) updateParent(node,
	parentOldChild, newParent *Node[T]) {
	node.parent = newParent
	if newParent != nil {
	if newParent.value > parentOldChild.value {
		newParent.left = node
	} else {
		newParent.right = node
	}
	} else {
	tree.root = node
	}
}

func search[T OrderedStringer](value T, node *Node[T]) (*Node[T], string) {
    if value == node.value {
        return nil, ""
    } else if value > node.value {
        if node.right == nil {
            return node, "R"
        }
        return search(value, node.right)
    } else if value < node.value {
        if node.left == nil {
            return node, "L"
        }
		return search(value, node.left)
    }
    return nil, ""
}

// Enhanced Search method that returns a boolean indicating presence
func (tree *RedBlackTree[T]) Search(value T) bool {
    return tree.IsPresent(value, tree.root)
}

// Delete removes a value from the tree
func (tree *RedBlackTree[T]) Delete(value T) {
    if tree.root == nil {
        return
    }
    
    // Find the node to delete
    node := tree.findNode(value)
    if node == nil {
        return // Value not found
    }
    
    tree.count--
    tree.deleteNode(node)
}

// Helper method to find a node with given value
func (tree *RedBlackTree[T]) findNode(value T) *Node[T] {
    current := tree.root
    for current != nil {
        if value == current.value {
            return current
        }
        if value < current.value {
            current = current.left
        } else {
            current = current.right
        }
    }
    return nil
}

// Main deletion logic
func (tree *RedBlackTree[T]) deleteNode(node *Node[T]) {
    var child *Node[T]
    
    // Case 1: Node has at most one child
    if node.left == nil || node.right == nil {
        if node.left != nil {
            child = node.left
        } else {
            child = node.right
        }
        
        // If deleting root
        if node.parent == nil {
            tree.root = child
            if child != nil {
                child.red = false
            }
            return
        }
        
        // Replace node with its child
        parent := node.parent
        if node == parent.left {
            parent.left = child
        } else {
            parent.right = child
        }
        if child != nil {
            child.parent = parent
        }
        
        // If deleted node was black, fix the tree
        if !node.red && (child == nil || !child.red) {
            tree.fixDelete(parent, child)
        }
        return
    }
    
    // Case 2: Node has two children
    // Find successor (smallest value in right subtree)
    successor := node.right
    for successor.left != nil {
        successor = successor.left
    }
    
    // Copy successor value to node
    node.value = successor.value
    // Delete the successor
    tree.deleteNode(successor)
}

// Fix tree after deletion of a black node
func (tree *RedBlackTree[T]) fixDelete(parent *Node[T], node *Node[T]) {
    for node != tree.root && (node == nil || !node.red) {
        var sibling *Node[T]
        isLeft := (node == parent.left || (node == nil && parent.left == nil))
        
        if isLeft {
            sibling = parent.right
            if sibling != nil && sibling.red {
                sibling.red = false
                parent.red = true
                tree.leftRotate(nil, sibling, parent, false)
                sibling = parent.right
            }
            
            if sibling == nil || (sibling.left == nil || !sibling.left.red) && 
                (sibling.right == nil || !sibling.right.red) {
                if sibling != nil {
                    sibling.red = true
                }
                if parent.red {
                    parent.red = false
                    break
                }
                node = parent
                parent = node.parent
                continue
            }
            
            if sibling.right == nil || !sibling.right.red {
                sibling.left.red = false
                sibling.red = true
                tree.rightRotate(nil, sibling.left, sibling, false)
                sibling = parent.right
            }
            
            sibling.red = parent.red
            parent.red = false
            sibling.right.red = false
            tree.leftRotate(nil, sibling, parent, false)
            break
        } else {
            sibling = parent.left
            if sibling != nil && sibling.red {
                sibling.red = false
                parent.red = true
                tree.rightRotate(nil, sibling, parent, false)
                sibling = parent.left
            }
            
            if sibling == nil || (sibling.left == nil || !sibling.left.red) && 
                (sibling.right == nil || !sibling.right.red) {
                if sibling != nil {
                    sibling.red = true
                }
                if parent.red {
                    parent.red = false
                    break
                }
                node = parent
                parent = node.parent
                continue
            }
            
            if sibling.left == nil || !sibling.left.red {
                sibling.right.red = false
                sibling.red = true
                tree.leftRotate(nil, sibling.right, sibling, false)
                sibling = parent.left
            }
            
            sibling.red = parent.red
            parent.red = false
            sibling.left.red = false
            tree.rightRotate(nil, sibling, parent, false)
            break
        }
    }
    
    if node != nil {
        node.red = false
    }
}
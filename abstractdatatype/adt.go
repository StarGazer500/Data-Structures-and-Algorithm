package abstractdatatype

import (
	"fmt"
)

// Stack implementation
type Stack[T any] struct {
	items []T
}

// Methods
func (stack *Stack[T]) Push(item T) {
	// item is added to the right-most position in the
	// slice
	stack.items = append(stack.items, item)
}
func (stack *Stack[T]) Pop() T {
	length := len(stack.items)
	returnValue := stack.items[length-1]
	stack.items = stack.items[:(length - 1)]
	return returnValue
}
func (stack Stack[T]) Top() T {
	length := len(stack.items)
	return stack.items[length-1]
}
func (stack Stack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

// Queue Implementation

type Queue[T any] struct {
	items []T
}
type Iterator[T any] struct {
	next  int // index in items
	items []T
}

// Queue Methods
//
//	func (queue *Queue[T]) Insert(item T) {
//	    // item is added to the right-most position in the slice
//	    queue.items = append(queue.items, item)
//	}
//
//	func (queue *Queue[T]) Remove() T {
//	    returnValue := queue.items[0]
//	    queue.items = queue.items[1:]
//	    return returnValue
//	}
func (queue Queue[T]) First() T {
	return queue.items[0]
}

func (qeue *Queue[T]) Last() T {
	return qeue.items[len(qeue.items)-1]
}

func (queue Queue[T]) Size() int {
	return len(queue.items)
}

func (queue *Queue[T]) Range() Iterator[T] {
	return Iterator[T]{0, queue.items}
}

// Iterator Methods
func (iterator *Iterator[T]) Empty() bool {
	return iterator.next == len(iterator.items)
}
func (iterator *Iterator[T]) Next() T {
	returnValue := iterator.items[iterator.next]
	iterator.next++
	return returnValue
}

func (queue *Queue[T]) InsertFront(item T) {
	queue.items = append(queue.items, item) // Expands deque.items
	for i := len(queue.items) - 1; i > 0; i-- {
		queue.items[i] = queue.items[i-1]
	}
	queue.items[0] = item
}

func (qeue *Queue[T]) InsertBack(item T) {
	qeue.items = append(qeue.items, item)
}

func (queue *Queue[T]) RemoveFirst() T {
	returnValue := queue.items[0]
	queue.items = queue.items[1:]
	return returnValue
}

func (qeue *Queue[T]) RemoveLast() T {
	length := len(qeue.items)
	returnValue := qeue.items[length-1]
	qeue.items = qeue.items[:(length - 1)]
	return returnValue
}

func (qeue *Queue[T]) Empty() bool {
	return len(qeue.items) == 0
}

// Priority Queue

type PriorityQueue[T any] struct {
	q    []Queue[T]
	size int
}

func NewPriorityQueue[T any](numberPriorities int) (pq PriorityQueue[T]) {
	pq.q = make([]Queue[T], numberPriorities)
	return pq
}

// Methods for priority queue
func (pq *PriorityQueue[T]) Insert(item T, priority int) {
	pq.q[priority-1].InsertBack(item)
	pq.size++
}

func (pq *PriorityQueue[T]) Remove() T {
	pq.size--
	for i := 0; i < len(pq.q); i++ {
		if pq.q[i].Size() > 0 {
			return pq.q[i].RemoveFirst()
		}
	}
	var zero T
	return zero
}

func (pq *PriorityQueue[T]) First() T {
	for _, queue := range pq.q {
		if queue.Size() > 0 {
			return queue.First()
		}
	}
	var zero T
	return zero
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	result := true
	for _, queue := range pq.q {
		if queue.Size() > 0 {
			result = false
			break
		}
	}
	return result
}

// SingleLinkedlist
type Ordered interface {
	~string | ~int | ~float64
}
type SingleLinkedListNode[T Ordered] struct {
	Item T
	next *SingleLinkedListNode[T]
}
type SingleLinkedList[T Ordered] struct {
	first       *SingleLinkedListNode[T]
	numberItems int
}

// Methods
func (list *SingleLinkedList[T]) Append(item T) {
	// Adds item to a new node at the end of the list
	newSingleLinkedListNode := SingleLinkedListNode[T]{item, nil}
	if list.first == nil {
		list.first = &newSingleLinkedListNode
	} else {
		last := list.first
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &newSingleLinkedListNode
	}
	list.numberItems += 1
}

func (list *SingleLinkedList[T]) InsertAt(index int, item T) error {
	// Adds item to a new node at position index in the list
	if index < 0 || index > list.numberItems {
		return fmt.Errorf("index out of bounds error")
	}
	newSingleLinkedListNode := SingleLinkedListNode[T]{item, nil}
	if index == 0 {
		newSingleLinkedListNode.next = list.first
		list.first = &newSingleLinkedListNode
		list.numberItems += 1
		return nil // No error
	}
	node := list.first
	count := 0
	previous := node
	for count < index {
		previous = node
		count++
		node = node.next
	}
	newSingleLinkedListNode.next = node
	previous.next = &newSingleLinkedListNode
	list.numberItems += 1
	return nil // no error
}

func (list *SingleLinkedList[T]) RemoveAt(index int) (T, error) {
	if index < 0 || index > list.numberItems {
		var zero T
		return zero, fmt.Errorf("Index out of bounds error")
	}
	node := list.first
	if index == 0 {
		toRemove := node
		list.first = toRemove.next
		list.numberItems -= 1
		return toRemove.Item, nil
	}
	count := 0
	previous := node
	for count < index {
		previous = node
		count++
		node = node.next
	}
	toRemove := node
	previous.next = toRemove.next
	list.numberItems -= 1
	return toRemove.Item, nil
}
func (list *SingleLinkedList[T]) IndexOf(item T) int {
	node := list.first
	count := 0
	for {
		if node.Item == item {
			return count
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		count += 1
	}
}

func (list *SingleLinkedList[T]) ItemAfter(item T) T {
	// Scan list for the first occurence of item
	node := list.first
	for {
		if node == nil { // item not found
			var zero T
			return zero
		}
		if node.Item == item {
			break
		}
		node = node.next
	}
	return node.next.Item
}
func (list *SingleLinkedList[T]) Items() []T {
	result := []T{}
	node := list.first
	for i := 0; i < list.numberItems; i++ {
		result = append(result, node.Item)
		node = node.next
	}
	return result
}
func (list *SingleLinkedList[T]) First() *SingleLinkedListNode[T] {
	return list.first
}
func (list *SingleLinkedList[T]) Size() int {
	return list.numberItems
}

// Double Linked DoubleLinkedList
type DoubleLinkedListNode[T Ordered] struct {
	Item T
	next *DoubleLinkedListNode[T]
	prev *DoubleLinkedListNode[T]
}
type DoubleLinkedList[T Ordered] struct {
	first       *DoubleLinkedListNode[T]
	last        *DoubleLinkedListNode[T]
	numberItems int
}

// Methods
func (list *DoubleLinkedList[T]) Append(item T) {
	// Adds item to a new node at the end of the list
	newNode := DoubleLinkedListNode[T]{item, nil, nil}
	if list.first == nil {
		list.first = &newNode
		list.last = list.first
	} else {
		list.last.next = &newNode
		newNode.prev = list.last
		list.last = &newNode
	}
	list.numberItems += 1
}
func (list *DoubleLinkedList[T]) InsertAt(index int, item T) error {
	// Adds item to a new node at position index in the list
	if index < 0 || index > list.numberItems {
		return fmt.Errorf("index out of bounds error")
	}
	newNode := DoubleLinkedListNode[T]{item, nil, nil}
	if index == 0 {
		newNode.next = list.first
		if list.first != nil {
			list.first.prev = &newNode
		}
		list.first = &newNode
		list.numberItems += 1
		if list.numberItems == 1 {
			list.last = list.first
		}
		return nil // No error
	}
	node := list.first
	count := 0
	previous := node
	for count < index {
		previous = node
		count++
		node = node.next
	}
	newNode.next = node
	previous.next = &newNode
	node.prev = &newNode
	newNode.prev = previous
	list.numberItems += 1
	return nil // no error
}
func (list *DoubleLinkedList[T]) RemoveAt(index int) (T, error) {
	if index < 0 || index > list.numberItems {
		var zero T
		return zero, fmt.Errorf("index out of bounds error")
	}
	node := list.first
	if index == 0 {
		toRemove := node
		list.first = toRemove.next
		list.numberItems -= 1
		if list.numberItems <= 1 {
			list.last = list.first
		}
		return toRemove.Item, nil
	}
	count := 0
	previous := node
	for count < index {
		previous = node
		count++
		node = node.next
	}
	toRemove := node
	previous.next = toRemove.next
	toRemove.next.prev = previous
	list.numberItems -= 1
	if list.numberItems <= 1 {
		list.last = list.first
	}
	return toRemove.Item, nil
}

func (list *DoubleLinkedList[T]) IndexOf(item T) int {
	node := list.first
	count := 0
	for {
		if node.Item == item {
			return count
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		count += 1
	}
}
func (list *DoubleLinkedList[T]) ItemAfter(item T) T {
	// Scan list for the first occurence of item
	node := list.first
	for {
		if node == nil { // item not found
			var zero T
			return zero
		}
		if node.Item == item {
			break
		}
		node = node.next
	}
	return node.next.Item
}
func (list *DoubleLinkedList[T]) ItemBefore(item T) T {
	// Scan list for the first occurence of item
	node := list.first
	for {
		if node == nil { // item not found
			var zero T
			return zero
		}
		if node.Item == item {
			break
		}
		node = node.next
	}
	return node.prev.Item
}
func (list *DoubleLinkedList[T]) Items() []T {
	result := []T{}
	node := list.first
	for i := 0; i < list.numberItems; i++ {
		result = append(result, node.Item)
		node = node.next
	}
	return result
}
func (list *DoubleLinkedList[T]) ReverseItems() []T {
	result := []T{}
	node := list.last
	for {
		if node == nil {
			break
		}
		result = append(result, node.Item)
		node = node.prev
	}
	return result
}

func (list *DoubleLinkedList[T]) First() *DoubleLinkedListNode[T] {
	return list.first
}
func (list *DoubleLinkedList[T]) Last() *DoubleLinkedListNode[T] {
	return list.last
}
func (list *DoubleLinkedList[T]) Size() int {
	return list.numberItems
}

package linearadt

import (
	"fmt"
)

// Stack Stack Test
func TestStackAbractDataType() {

	nameStack := Stack[string]{}
	nameStack.Push("Zachary")
	nameStack.Push("Adolf")
	if !nameStack.IsEmpty() {
		topOfStack := nameStack.Top()
		fmt.Printf("\nTop of stack is %s", topOfStack)
	}
	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s",
			poppedFromStack)
	}
	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s",
			poppedFromStack)
	}
	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s",
			poppedFromStack)
	}
	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s",
			poppedFromStack)
	}
	// Create a stack of integers
	intStack := Stack[int]{}
	intStack.Push(5)
	intStack.Push(10)
	intStack.Push(0)

	if !intStack.IsEmpty() {
		top := intStack.Top()
		fmt.Printf("\nValue on top of intStack is %d", top)
	}
	if !intStack.IsEmpty() {
		popFromStack := intStack.Pop()
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}
	if !intStack.IsEmpty() {
		popFromStack := intStack.Pop()
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}
	if !intStack.IsEmpty() {
		popFromStack := intStack.Pop()
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}
}

// Queue Test
func TestQueueAbstractDatatype() {
	myQueue := Queue[int]{}
	myQueue.InsertBack(15)
	myQueue.InsertBack(20)
	myQueue.InsertBack(30)
	myQueue.RemoveFirst()
	fmt.Println(myQueue.First())
	queue := Queue[float64]{}
	for i := 0; i < 10; i++ {
		queue.InsertBack(float64(i))
	}
	iterator := queue.Range()
	for {
		if iterator.Empty() {
			break
		}
		fmt.Println(iterator.Next())
	}
	fmt.Println("queue.First() = ", queue.First())

	myQueue.InsertFront(5)
	myQueue.InsertBack(10)
	myQueue.InsertFront(2)
	myQueue.InsertBack(12) // 2 5 10 12
	fmt.Println("myDeque.First() = ", myQueue.First())
	fmt.Println("myDeque.Last() = ", myQueue.Last())
	myQueue.RemoveLast()
	myQueue.RemoveFirst()
	fmt.Println("myDeque.First() = ", myQueue.First())
	fmt.Println("myDeque.Last() = ", myQueue.Last())
}

// Testing Priority Queue

type Passenger struct {
	name     string
	priority int
}

func TestPriorityQueueAbstractDatatype() {

	airlineQueue := NewPriorityQueue[Passenger](3)
	passengers := []Passenger{{"Erika", 3}, {"Robert", 3}, {"Danielle", 3},

		{"Madison", 1}, {"Frederik", 1}, {"James", 2},
		{"Dante", 2}, {"Shelley", 3}}
	fmt.Println("Passsengers: ", passengers)
	for i := 0; i < len(passengers); i++ {
		airlineQueue.Insert(passengers[i], passengers[i].priority)
	}
	fmt.Println("First passenger in line: ", airlineQueue.First())
	airlineQueue.Remove()
	airlineQueue.Remove()
	airlineQueue.Remove()
	fmt.Println("First passenger in line: ", airlineQueue.First())
}

// SingleLinkedlist Test
func SingleLinkedlist() {
	cars := SingleLinkedList[string]{}
	cars.Append("Honda")
	cars.InsertAt(0, "Nissan")
	cars.InsertAt(0, "Chevy")
	cars.InsertAt(1, "Ford")
	cars.InsertAt(1, "Tesla")
	cars.InsertAt(0, "Audi")
	cars.InsertAt(2, "Volkswagon")
	cars.Append("Volvo")
	fmt.Println(cars.Items())
	fmt.Println("Index of Tesla: ", cars.IndexOf("Tesla"))
	cars.RemoveAt(0)
	car, _ := cars.RemoveAt(3)
	fmt.Println("car removed is: ", car)
	fmt.Println(cars.Items())
	cars.RemoveAt(cars.Size() - 1)
	fmt.Println(cars.Items())
	cars.Append("Lexus")
	fmt.Println(cars.Items())
	fmt.Println("First car in the list is: ", cars.First().Item)
	fmt.Println("Last car in the list is: ", cars.Items()[cars.Size()-1])
}

func TestHeapPriorityQueue(){
	myQueue := HeapPriorityQueue[string]{}
    myQueue.Push("Helen")
    myQueue.Push("Apollo")
    myQueue.Push("Richard")
    myQueue.Push("Barbara")
    fmt.Println(myQueue)
    myQueue.Pop()
    fmt.Println(myQueue)
    myQueue.Push("Arlene")
    fmt.Println(myQueue)
    myQueue.Pop()
    myQueue.Pop()
    fmt.Println(myQueue)
}



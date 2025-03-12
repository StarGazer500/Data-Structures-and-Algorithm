package heap


import("fmt")

func TestHeapTree(){
	slice1:= []int{100, 60, 80, 50, 30, 75, 40, 10, 35}
    heap1 := NewHeap[int](slice1)
    heap1.Insert(90)
    fmt.Println("heap1 after inserting 90")
    fmt.Println(heap1.Items)
    fmt.Println("Largest item in heap: ", heap1.Largest())
    heap1.Remove()
    fmt.Println("Removing largest item from heap yielding the heap: ")
    fmt.Println(heap1.Items)
    fmt.Println("Largest item in heap: ", heap1.Largest())
    slice2:= []int{10, 35, 100, 80, 30, 75, 40, 50, 60}
    heap2 := NewHeap[int](slice2)
    heap2.Insert(90)
    fmt.Println("heap2 with rearranged slice2 after inserting 90")
    fmt.Println(heap2.Items)
}
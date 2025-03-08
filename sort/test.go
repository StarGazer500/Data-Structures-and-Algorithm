package sort

import ("fmt"
       "time"
	   "math/rand"
	    
)

func Testbublesort(){
	numbers := []float64{3.5, -2.4, 12.8, 9.1}
    names := []string{"Zachary", "John", "Moe", "Jim", "Robert"}
    Bubblesort[float64](numbers)
    fmt.Println(numbers)
    Bubblesort[string](names)
    fmt.Println(names)

}


func TestInsertsort(){
	numbers := []float64{3.5, -2.4, 12.8, 9.1}
    names := []string{"Zachary", "John", "Moe", "Jim", "Robert"}
    InsertSort[float64](numbers)
    fmt.Println(numbers)
    InsertSort[string](names)
    fmt.Println(names)

}

const size = 50_000_000
func TestConcurrentQuickSort(){
    
	data := make([]float64, size)
    for i := 0; i < size; i++ {
        data[i] = 100.0 * rand.Float64()
    }
    data2 := make([]float64, size)
    copy(data2, data)
    start := time.Now()
    QSort[float64](data)
    elapsed := time.Since(start)
    fmt.Println("Elapsed time for concurrent quicksort = ", elapsed)
}


func TestNonConcurrentQuickSort(){
   
	data := make([]float64, size)
    for i := 0; i < size; i++ {
        data[i] = 100.0 * rand.Float64()
    }
    data2 := make([]float64, size)
  
    start := time.Now()
    
    Quicksort(data2, 0, len(data2) - 1)
    elapsed := time.Since(start)
    fmt.Println("Elapsed time for regular quicksort = ", elapsed)
    fmt.Println("Is sorted: ", IsSorted(data2))
}


func TestConcurrentMergeSort(){
    data := make([]float64, size)
    for i := 0; i < size; i++ {
        data[i] = 100.0 * rand.Float64()
    }
    start := time.Now()
    result := ConcurrentMergeSort(data)
    elapsed := time.Since(start)
    fmt.Println("Elapsed time for concurrent mergesort = ", elapsed)
    fmt.Println("Sorted: ", IsSorted(result))
}



package linearsearch

import(
	"fmt"
    "time"
    "math/rand"
)
const size = 100_000_000

func TestBruteForceLinearSearch(){
	data := make([]float64, size)
    for i := 0; i < size; i++ {
        data[i] = 100.0 * rand.Float64()
    }
    start := time.Now()
    result := ConcurrentBruteForceSearch[float64](data, 54.0) // Should return false
    elapsed := time.Since(start)
    fmt.Println("Time to search slice of 100_000_000 floats using concurrentSearch = ", elapsed)
    fmt.Println("Result of search is ", result)

	start = time.Now()
    result = ConcurrentBruteForceSearch[float64](data, data[size / 2]) // true
    elapsed = time.Since(start)
    fmt.Println("Time to search slice of 100_000_000 floats usingconcurrentSearch = ", elapsed)
    fmt.Println("Result of search is ", result)
}

func TestBinarySearch(){
	data := make([]float64, size)
    for i := 0; i < size; i++ {
        data[i] = float64(i) // is sorted
    }
    start := time.Now()
    result := BinarySearch[float64](data, -10.0)
    elapsed := time.Since(start)
    fmt.Println("Time to search slice of 100_000_000 floats using binarySearch = ", elapsed)
    fmt.Println("Result of search is ", result)
	
    start = time.Now()
    result = BinarySearch[float64](data, float64(size / 2))
    elapsed = time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using binarySearch = ", elapsed)
    fmt.Println("Result of search is ", result)
}


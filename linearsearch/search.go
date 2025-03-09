package linearsearch
import "runtime"


// Concurrent Brute Force Search
type Ordered interface {
    ~float64 | ~int | ~string
}


func searchSegment[T Ordered](slice []T, target T, a, b int, ch chan<- bool) {
		// Generates boolean value put into ch
		for i := a; i < b; i++ {
			if slice[i] == target {
				ch <- true
			}
		}
		ch <- false
	}

func ConcurrentBruteForceSearch[T Ordered](data []T, target T) bool {
    ch := make(chan bool)
    numSegments := runtime.NumCPU()
    segmentSize := int(float64(len(data)) / float64(numSegments))
    // Launch numSegments goroutines
    for index := 0; index < numSegments; index++ {
        go searchSegment(data, target, index * segmentSize, index *
                             segmentSize + segmentSize, ch)
    }
    num := 0 // Completed goroutines
    for {
        select {
        case value := <- ch:  // Blocks until a goroutine puts a bool into the
                              //channel
            if value == true {
                return true
            }
            num += 1

            if num == numSegments { // All goroutiines have completed
                return false
            }
        }
    }
    return false
}


func BinarySearch[T Ordered](slice []T, target T) bool {
    low := 0
    high := len(slice) - 1
    for low <= high {
        median := (low + high) / 2
        if slice[median] < target {
            low = median + 1
        } else {
            high = median - 1
        }
    }

    if low == len(slice) || slice[low] != target {
        return false
    }
    return true
}
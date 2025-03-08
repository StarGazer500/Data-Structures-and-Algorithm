package sort
import(
    "sync"
)

type Ordered interface {
    ~float64 | ~int | ~string
}

func Bubblesort[T Ordered](data []T) {
    n := len(data)
    for i:= 0; i < n - 1; i++ {
        for j:= 0; j < n - 1 - i; j++ {
            if data[j] > data[j + 1] {


                data[j], data[j + 1] = data[j + 1], data[j]
            }
        }
    }
}

func InsertSort[T Ordered](data[] T) {
    i := 1
    for i < len(data) {
        h := data[i]
        j := i - 1
        for j >= 0 && h < data[j] {
            data[j + 1] = data[j]
            j -= 1
        }
        data[j + 1] = h
        i += 1
    }
}

// Noncuncrrent Quick Sort Algorithm
func partition[T Ordered](data []T, low, high int) int {
    var pivot = data[low]
    var i = low
    var j = high
    for i < j {
        for data[i] <= pivot && i < high {
            i++;
        }
        for data[j] > pivot && j > low {
            j--
        }
        if i < j {  data[i], data[j] = data[j], data[i]
        }
    }
    data[low] = data[j]
    data[j] = pivot
    return j
}


func Quicksort[T Ordered](data []T, low, high int) {
    if low < high {
        var pivot = partition(data, low, high)
        Quicksort(data, low, pivot)
        Quicksort(data, pivot + 1, high)
    }
}




const threshold = 5000

func Partition[T Ordered](data[] T) int {
    data[len(data) / 2], data[0] = data[0], data[len(data) / 2]
    pivot := data[0]
    mid := 0
    i := 1
    for i < len(data) {
        if data[i] < pivot {
            mid += 1
            data[i], data[mid] = data[mid], data[i]
        }
        i += 1
    }
    data[0], data[mid] = data[mid], data[0]
    return mid
}

func ConcurrentQuicksort[T Ordered](data[] T, wg *sync.WaitGroup) {
    for len(data) >= 30 {
        mid := Partition(data)
        var portion[] T
        if mid < len(data) / 2 {
            portion = data[:mid]
            data = data[mid + 1:]
            } else {
                portion = data[mid + 1:]
                data = data[:mid]
            }
            if (len(portion) > threshold) {
                wg.Add(1)
                go func(data[] T) {
                    defer wg.Done()
                    ConcurrentQuicksort(data, wg)
                }(portion)
            } else {
                ConcurrentQuicksort(portion, wg)
            }
        }
        InsertSort(data)
    }

    func QSort[T Ordered](data[] T) {
        var wg sync.WaitGroup
        ConcurrentQuicksort(data, &wg)
        wg.Wait()
    }

    func IsSorted[T Ordered](data[] T) bool {
        for i := 1; i < len(data); i++ {
            if data[i] < data[i - 1] {
                return false
            }
        }
        return true
    }

    // Concurrent Mergesort
    const max = 5000
    func Merge[T Ordered](left, right []T) []T {
        result := make([]T, len(left) + len(right))
        i, j, k := 0, 0, 0
   
        for i < len(left) && j < len(right) {
            if left[i] < right[j] {
                result[k] = left[i]
                i++
            } else {
                result[k] = right[j]
                j++
            }
            k++
        }
        for i < len(left) {
            result[k] = left[i]
            i++
            k++
        }
        for j < len(right) {
            result[k] = right[j]
            j++
            k++
        }
        return result
    }

    func MergeSort[T Ordered](data []T) []T {
        if len(data) > 100 {
            middle := len(data) / 2
            left := data[:middle]
            right := data[middle:]
            data = Merge(MergeSort(right), MergeSort(left))
        } else {
            InsertSort(data)
        }
        return data
    }

    func ConcurrentMergeSort[T Ordered](data []T) []T {
        if len(data) > 1 {
            if len(data) <= max {
                return MergeSort(data)
            } else { // Concurrent
                middle := len(data) / 2
                left := data[:middle]
                right := data[middle:]
                var wg sync.WaitGroup
                wg.Add(2)
                var data1, data2 []T
                go func() {
                    defer wg.Done()
                    data1 = ConcurrentMergeSort(left)
                }()
                go func() {
                    defer wg.Done()
                    data2 = ConcurrentMergeSort(right)
                }()
                wg.Wait()
                return Merge(data1, data2)
            }
        }
        return nil
    }
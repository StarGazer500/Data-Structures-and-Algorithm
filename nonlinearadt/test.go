package nonlinearadt

import (
    "fmt"
     // Fowler–Noll–Vo algorithm
    "strconv"
    "time"
)



func TestHashTable(){
	myTable := NewTable()
mapCollection := make(map[string]string)
words := []string{}
for i := 0; i < 500_000; i++ {
	word := strconv.Itoa(i)
	words = append(words, word)
	myTable.Insert(word)
	mapCollection[word] = ""
}
fmt.Println("Benchmark test begins to test words: ", length)
start := time.Now()
for i := 0; i < length; i++ {
	if !myTable.IsPresent(words[i]) {
		fmt.Println("Word not found in table: ", words[i])
	}
}
elapsed := time.Since(start)

fmt.Println("Time to test all words in myTable: ", elapsed)
start = time.Now()
for i := 0; i < len(mapCollection); i++ {
	_, present := mapCollection[words[i]]
	if !present {
		fmt.Println("Word not found in mapCollection: ", words[i])
	}
}
elapsed = time.Since(start)
fmt.Println("Time to test words in mapCollection: ", elapsed)
}

func PatternSearch(){
	text :="31415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"
    pattern := "816406286208998628034825342"
    start := time.Now()
    _, _ = BruteForceSearch(text, pattern)
    elapsed := time.Since(start)
    fmt.Println("Computation time using BruteForceSearch: ", elapsed)
    start = time.Now()
    _, _ = Search(text, pattern)
    elapsed = time.Since(start)
    fmt.Println("Computation time using Search: ", elapsed)
    fmt.Println(BruteForceSearch(text, pattern))
    fmt.Println(Search(text, pattern))
}


func TestSet(){
	set1 := Set[int]{}
    set1.Insert(3)
    set1.Insert(5)
    set1.Insert(7)
    set1.Insert(9)
    set2 := Set[int]{}
    set2.Insert(3)
    set2.Insert(6)
    set2.Insert(8)
    set2.Insert(9)
    set2.Insert(11)
    set2.Delete(11)
    fmt.Println("Items in set2: ", set2.Items())
    fmt.Println("5 in set1: ", set1.In(5))
    fmt.Println("5 in set2: ", set2.In(5))
    fmt.Println("Union of set1 and set2: ", set1.Union(set2).Items())
    fmt.Println("Intersection of set1 and set2: ",
                     set1.Intersection(set2).Items())
    fmt.Println("Difference of set2 with respect to set1: ",
                     set2.Difference(set1).Items())
    fmt.Println("Size of this difference: ", set1.
Intersection(set2).Size())
}
package nonlinearadt

import(
	"hash/fnv"
	 "strings"
)

// Hashtable Implementation
const tableSize = 100_000
var length int

func hash(s string) uint32 {
    h := fnv.New32a() // Fowler-Noll-Vo algorithm
    h.Write([]byte(s))
    return h.Sum32()
}
type WordType struct {
    word string
    list []string
}

type HashTable [tableSize]WordType

func NewTable() HashTable {
    var table HashTable
    for i := 0; i < tableSize; i++ {
        table[i] = WordType{"", []string{}}
    }
    return table
}


func (table *HashTable) Insert(word string) {
    index := hash(word) % tableSize // Between 0 and tableSize - 1
    // Search table[index] for word
    if table[index].word == word {
        return // duplicates not allowed
    }
    if len(table[index].list) > 0 {
        for i := 0; i < len(table[index].list); i++ {
            if table[index].list[i] == word {
                return // duplicates not allowed
            }
        }
    }
    if table[index].word == "" {
        table[index].word = word
    } else {
        table[index].list = append(table[index].list, word)
    }
    length += 1
}


func (table HashTable) IsPresent(word string) bool {
    index := hash(word) % tableSize // Between 0 and tableSize - 1
    // Search table[index] for word
    if table[index].word == word {
        return true
    }
    if len(table[index].list) > 0 {
		for i := 0; i < len(table[index].list); i++ {
            if table[index].list[i] == word {
                return true
            }
        }
    }
    return false
}

// Pattern Search 

const (
    Radix = uint64(10)
    Q     = uint64(10 ^ 9 + 9)
)
func Hash(s string, Length int) uint64 {
    // Horner's method
    h := uint64(0)
    for i := 0; i < Length; i++ {
        h = (h*Radix + uint64(s[i])) % Q
    }
    return h
}

func Search(txt, pattern string) (bool, int) {
    strings.ToLower(txt)
    strings.ToLower(pattern)
    n := len(txt)
    m := len(pattern)
	patternHash := Hash(pattern, m)
    textHash := Hash(txt, m)
    if textHash == patternHash {
        return true, 0
    }
    PM := uint64(1)
    for i := 1; i <= m-1; i++ {
        PM = (Radix * PM) % Q
    }
    for i := m; i < n; i++ {
        textHash = (textHash + Q - PM*uint64(txt[i-m])%Q) % Q
        textHash = (textHash*Radix + uint64(txt[i])) % Q
        if (patternHash == textHash) && pattern == txt[(i-m+1):(i+1)] {
            return true, i - m + 1
        }
    }
    return false, -1
}

func BruteForceSearch(txt, pattern string) (bool, int) {
    patternLength := len(pattern)
    for outer := 0; outer < len(txt)-patternLength; outer++ {
        if txt[(outer):(outer+patternLength)] == pattern {
            return true, outer
        }
    }
    return false, -1
}

// Set Implementation
type Ordered interface {
    ~string | ~int | ~float64
}
type Set[T Ordered] struct {
    items map[T]bool
}

func (set *Set[T]) Insert(item T) {
    if set.items == nil {
        set.items = make(map[T]bool)
    }
    // Prevent duplicate entry
    _, present := set.items[item]
    if !present {
        set.items[item] = true
    }
}

func (set *Set[T]) Delete(item T) {
    _, present := set.items[item]
	if present {
        delete(set.items, item)
    }
}

func (set *Set[T]) In(item T) bool {
    _, present := set.items[item]
    return present
}

func (set *Set[T]) Items() []T {
    items := []T{}
    for item := range set.items {
        items = append(items, item)
    }
    return items
}

func (set *Set[T]) Size() int {
    return len(set.items)
}

func (set *Set[T]) Union(set2 Set[T]) *Set[T] {
    result := Set[T]{}
    result.items = make(map[T]bool)
    for index := range set.items {
        result.items[index] = true
    }
    for j := range set2.items {
        _, present := result.items[j]
        if !present {
            result.items[j] = true
        }
    }
    return &result
}

func (set *Set[T]) Intersection(set2 Set[T]) *Set[T] {
    result := Set[T]{}
    result.items = make(map[T]bool)
    for i := range set2.items {
        _, present := set.items[i]
        if present {
            result.items[i] = true
        }
    }
	return &result
}
// Return a new set of items in set not found in set2
func (set *Set[T]) Difference(set2 Set[T]) *Set[T] {
    result := Set[T]{}
    result.items = make(map[T]bool)
    for i := range set.items {
        _, present := set2.items[i]
        if !present {
            result.items[i] = true
        }
    }
    return &result
}

func (set *Set[T]) Subset(set2 Set[T]) bool {
    for i := range set.items {
        _, present := set2.items[i]
        if !present {
            return false
        }
    }
    return true
}
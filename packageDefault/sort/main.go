package main

import (
	"fmt"
	"sort"
)

// sort

type Entry struct {
	Name  string
	Value int
}

func main() {
	i := []int{5, 3, 5, 6, 4, 4, 3, 56, 77, 23, 4, 234, 5, 234}
	s := []string{"a", "s", "f", "g"}

	fmt.Println("before sort", i, s) //[5 3 5 6 4 4 3 56 77 23 4 234 5 234] [a s f g]

	sort.Ints(i)
	sort.Strings(s)

	fmt.Println("after sort", i, s) //[3 3 4 4 4 5 5 5 6 23 56 77 234 234] [a f g s]

}

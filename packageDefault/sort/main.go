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

	// slice,sliceStable
	el := []Entry{
		{"A", 20},
		{"F", 30},
		{"e", 50},
		{"g", 10},
		{"c", 550},
		{"e", 640},
		{"h", 1230},
		{"Z", 20},
		{"E", 60},
	}
	fmt.Println(el)

	// slice
	sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name }) //名前の昇順にソートする
	fmt.Println("after sort by Name", el)
	sort.Slice(el, func(i, j int) bool { return el[i].Value < el[j].Value })
	fmt.Println("after sort by Value", el)

	// sliceStable
	sort.SliceStable(el, func(i, j int) bool { return el[i].Name < el[j].Name })
	fmt.Println("after sort with sliceStable", el)
	sort.SliceStable(el, func(i, j int) bool { return el[i].Value < el[j].Value })
	fmt.Println("after sort with sliceStable", el)

	//slice,sliceStableの違いは同じkeyやvalueがあった場合のsortの挙動をどうするのか
}

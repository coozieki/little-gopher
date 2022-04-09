package main

import (
	"fmt"
	"linked-list/list"
)

func main() {
	l := list.FromArray[bool]([]bool{true, false, true})
	_ = l.Delete(1)
	_ = l.Delete(1)
	_ = l.Insert(1, false)
	val, _ := l.At(1)
	fmt.Println(val)
	fmt.Println(l.Search(false))
}

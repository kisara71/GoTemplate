package main

import (
	"GoTemplate/Slice"
	"fmt"
)

func main() {
	slice := make([]int, 0)
	slice = append(slice, 1)
	slice, err := Slice.Insert[int](slice, 2, 1)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(slice)
	slice, err = Slice.Delete[int](slice, 1)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(slice)
}

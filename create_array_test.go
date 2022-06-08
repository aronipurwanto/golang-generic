package golang_generic

import (
	"fmt"
	"testing"
)

func NewArray[T any](b, k int) [][]T {
	var array = make([][]T, b)
	for i := range array {
		array[i] = make([]T, k)
	}
	return array
}

func Kotak(n int) [][]int {
	array := NewArray[int](n, n)
	angka := 1
	for i := 0; i < 5; i++ {
		for k := 0; k < 5; k++ {
			array[i][k] = angka
			angka += 5
		}
	}
	return array
}

func TestArray(t *testing.T) {
	arr := Kotak(5)
	fmt.Println(arr)
	fmt.Println()
}

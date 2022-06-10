package golang_generic

import (
	"fmt"
	"testing"
)

// type baru namanya Bag, sebagai Array
type Bag[T any] []T

// function yang digunakan untuk print type Bag
func PrintBag[T any](bag Bag[T]) {
	// karena bag adalah array, maka untuk print nya harus dengan for
	for _, value := range bag {
		fmt.Print(value, "\t")
	}
	fmt.Println()
}

func TestBag(t *testing.T) {
	number := Bag[int]{1, 2, 3, 4, 5}
	PrintBag(number)

	names := Bag[string]{"Rio", "Ari", "Bambang", "Aji", "Joko"}
	PrintBag(names)
}

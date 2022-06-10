package golang_generic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data[T any] struct {
	FirstName T
	LastName  T
}

func (d *Data[_]) SayHallo(name string) string {
	return "Hello " + name
}

func (d *Data[T]) changeFistName(firstName T) T {
	d.FirstName = firstName
	return firstName
}

func TestMethodData(t *testing.T) {
	data := Data[string]{
		FirstName: "Roni",
		LastName:  "Purwanto",
	}

	// testing untuk test changeName
	changeName := data.changeFistName("Rino")
	assert.Equal(t, "Rino", changeName)

	// testing untuk sayhalleo
	hello := data.SayHallo("Roni")
	assert.Equal(t, "Hello Roni", hello)
}

func TestData(t *testing.T) {
	data := Data[string]{
		FirstName: "Ahmad",
		LastName:  "Roni",
	}

	fmt.Println(data)

	data2 := Data[int]{
		FirstName: 10,
		LastName:  200,
	}
	fmt.Println(data2)
}

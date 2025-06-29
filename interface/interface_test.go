package _interface

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Animal interface {
	Say() (string, error)
}
type Duck struct{}

func (duck Duck) Say() (string, error) {
	return "Krya", nil
}

func animalSay(animal Animal) {
	fmt.Println(animal.Say())
}

func TestInterfaces(t *testing.T) {
	duck := Duck{}
	animalSay(duck)
	var _ Animal = (*Duck)(nil)
}

func TestInterfaceNil(t *testing.T) {
	var i interface{}
	assert.Equal(t, true, i == nil)
	var d Duck = i.(Duck)
	assert.Equal(t, true, d == nil)
}

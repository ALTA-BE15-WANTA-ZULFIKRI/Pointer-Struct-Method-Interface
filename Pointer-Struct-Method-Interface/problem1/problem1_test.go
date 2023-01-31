package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func TestSwap(t *testing.T) {
	a := 10
	b := 20

	swap(&a, &b)

	assert.Equal(t, a, 20, "a should be 20")
	assert.Equal(t, b, 10, "b should be 10")
}

func main() {
	fmt.Println("Test Swap")
	testing.RunTests(func(name string, t *testing.T) {
		t.Run()
	})
}
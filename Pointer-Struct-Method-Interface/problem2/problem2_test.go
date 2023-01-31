ppackage main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getMinMax(numbers ...*int) (min int, max int) { 
	min = *numbers[0]
	max = *numbers[0]
	for _, number := range numbers {
		if *number < min {
			min = *number
		}
		if *number > max {
			max = *number
		}
  }

  return min, max
}

func TestGetMinMax(t *testing.T) {
	a1 := 1
	a2 := 2
	a3 := 3
	a4 := 4
	a5 := 5
	a6 := 6

	min, max := getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	assert.Equal(t, min, 1)
	assert.Equal(t, max, 6)
}

func main() {
	fmt.Println("Test berhasil")
}
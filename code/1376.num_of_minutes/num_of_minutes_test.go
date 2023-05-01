package code

import (
	"fmt"
	"testing"
)

func TestNumOfMinutes(t *testing.T) {
	fmt.Println(numOfMinutes(15, 0, []int{-1, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}, []int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}))
}

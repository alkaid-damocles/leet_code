package code

import (
	"fmt"
	"testing"
)

func TestMakeArrayIncreasing(t *testing.T) {
	fmt.Println(makeArrayIncreasing([]int{1, 5, 3, 10000, 7, 200000000, 200000001}, []int{1, 3, 2, 2000000}))
}

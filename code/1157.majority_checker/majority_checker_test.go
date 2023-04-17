package code

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	mc := Constructor([]int{1, 1, 2, 2, 1, 1})
	fmt.Println(mc.Query(2, 3, 2))
}

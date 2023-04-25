package code

import (
	"fmt"
	"testing"
)

func TestSortPeople(t *testing.T) {
	fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
}

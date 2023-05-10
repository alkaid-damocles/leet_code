package code

import (
	"fmt"
	"testing"
)

func TestSmallestRepunitDivByK(t *testing.T) {
	a := 10

	b := a % 3
	c := (a*10 + 1) % 3
	d := ((a%3)*10 + 1) % 3
	fmt.Println(a, b, c, d)
	fmt.Println(smallestRepunitDivByK(3))
}

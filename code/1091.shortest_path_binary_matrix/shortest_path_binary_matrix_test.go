package code

import (
	"fmt"
	"testing"
)

func TestShortestPathBinaryMatrix(t *testing.T) {
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 1}}))
}

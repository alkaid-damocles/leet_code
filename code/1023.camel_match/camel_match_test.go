package code

import (
	"fmt"
	"testing"
)

func TestCamelMatch(t *testing.T) {
	fmt.Println(camelMatch([]string{"CompetitiveProgramming", "CounterPick", "ControlPanel"}, "ControlPanel"))
}

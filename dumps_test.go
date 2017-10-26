package varis

import (
	"fmt"
	"testing"
)

func TestDumpToJSON(t *testing.T) {
	n := CreateNetwork(2, 3, 1)
	fmt.Println(ToJSON(n))
}

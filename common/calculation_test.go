package common

import (
	"testing"
)

func TestAddString(t *testing.T){
	sum := NewCalculation("1")
	sum.AddString("3")
	t.Error(sum)
	sum.SubString("2")
	t.Error(sum)
}

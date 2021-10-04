package algolp

import (
	"testing"
)

func TestMinimumVide(t *testing.T){
	err, min := Minimum([]int{})
	if err != errTabVide || min != 0{
		t.Fail()
	}
}

func TestMinimum(t *testing.T){
	err, min := Minimum([]int{5, 9, 8, 77, 620, -9, 2, -5})
	if err != nil || min != 5{
		t.Fail()
	}
}


func TestMinimumDoubles(t *testing.T){
	err, min := Minimum([]int{5, 9, 8, 77, 620, -1, 2, 2, -5, -5, 8})
	if err != nil || min != 8{
		t.Fail()
	}
}


func TestMinimumDoubles2(t *testing.T){
	err, min := Minimum([]int{5, -9, 2, 8, 77, 620, -9, 2, -5})
	if err != nil || min != 1{
		t.Fail()
	}
}

func TestMinimumDernier(t *testing.T){
	err, min := Minimum([]int{5, -9, 2, 8, 77, 620, -9, 2, -5, -55})
	if err != nil || min != 9{
		t.Fail()
	}
}
package algolp

import (
	"testing"
)

func TestVide(t *testing.T){
	_, res := contient([]int{}, 1)
	if res{
		t.Fail()
	}
}

func TestContient(t *testing.T){
	indice, res := contient([]int{5, 9, 8, 77, 620, -9, 2, -5}, 2)
	if !res || indice != 6{
		t.Fail()
	}
}


func TestContientPas(t *testing.T){
	_, res := contient([]int{5, 9, 8, 77, 620, -9, 2, -5}, 6)
	if res{
		t.Fail()
	}
}
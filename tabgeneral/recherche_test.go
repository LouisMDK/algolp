package algolp

import (
	"testing"
)

func TestRechercheVide(t *testing.T){
	indice, res := Recherche([]int{}, 1)
	if res || indice != -1{
		t.Fail()
	}
}

func TestRecherche(t *testing.T){
	indice, res := Recherche([]int{5, 9, 8, 77, 620, -9, 2, -5}, 2)
	if !res || indice != 6{
		t.Fail()
	}
}


func TestRechercheDoubles(t *testing.T){
	indice, res := Recherche([]int{5, 9, 8, 77, 620, -9, 2, 2, -5}, 2)
	if !res || indice != 6{
		t.Fail()
	}
}


func TestRechercheDoubles2(t *testing.T){
	indice, res := Recherche([]int{5, 9, 2, 8, 77, 620, -9, 2, -5}, 2)
	if !res || indice != 2{
		t.Fail()
	}
}



func TestRecherchePas(t *testing.T){
	indice, res := Recherche([]int{5, 9, 8, 77, 620, -9, 2, -5}, 6)
	if res || indice != -1{
		t.Fail()
	}
}
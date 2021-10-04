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
	indice, res := Recherche([]int{-9, -5, 2, 5, 8, 9, 77, 620}, 2)
	if !res || indice == -1{
		t.Fail()
	}
}


func TestRechercheDoubles(t *testing.T){
	indice, res := Recherche([]int{-9, -5, 2, 2, 5, 8, 9, 77, 620}, 2)
	if !res || indice == -1{
		t.Fail()
	}
}


func TestRechercheDoubles2(t *testing.T){
	indice, res := Recherche([]int{-9, 2, 2, 2, 2, 9, 77, 620}, 2)
	if !res || indice == -1{
		t.Fail()
	}
}



func TestRecherchePas(t *testing.T){
	indice, res := Recherche([]int{-9, -5, 2, 5, 8, 9, 77, 620}, -4)
	if res || indice != -1{
		t.Fail()
	}
}
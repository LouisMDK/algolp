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

func BenchmarkRecherche(b *testing.B){
	for n:=0; n<b.N;n++{
		Recherche([]int{-9, -5, 2, 2, 3, 4, 5, 5, 6, 7, 8, 8, 8, 9, 15, 25, 33, 44, 55, 66, 68, 77, 100, 120, 156, 187, 205, 354, 486, 514, 620, 705, 805, 906, 1050, 1060, 1070, 1080, 2050, 3060, 9999, 19200}, 354)
	}
}
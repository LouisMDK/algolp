package algolp

import (
	"testing"
)

func TestPlaceInsertionVide(t *testing.T){
	var liste []int
	TriInsertionPlace(liste)
	if len(liste) != 0 || liste != nil{
		t.Fail()
	}
}

func TestPlaceInsertionVide2(t *testing.T){
	var liste []int = nil
	TriInsertionPlace(liste)
	if len(liste) != 0 || liste != nil {
		t.Fail()
	}
}

func TestPlaceDejaTrie(t *testing.T) {
	var liste []int = []int{1, 2, 36, 72, 80}
	TriInsertionPlace(liste)
	if !listeEgal(liste, []int{1, 2, 36, 72, 80}) {
		t.Fail()
	}
}

func TestPlaceValeursNegatives(t *testing.T) {
	var liste []int = []int{-5, -50, -63, -05, -98}
	TriInsertionPlace(liste)
	if !listeEgal(liste, []int{-98, -63, -50, -5, -5}) {
		t.Fail()
	}
}


func TestPlaceUneValeur(t *testing.T) {
	var liste []int = []int{1}
	TriInsertionPlace(liste)
	if !listeEgal(liste, []int{1}) {
		t.Fail()
	}
}

func TestPlacePlusieursValeurs(t *testing.T) {
	var liste []int = []int{9, 5, 69, 1, 0}
	TriInsertionPlace(liste)
	if !listeEgal(liste, []int{0, 1, 5, 9, 69}) {
		t.Fail()
	}
}

func TestPlaceDoublons(t *testing.T) {
	var liste []int = []int{-5, 2, 30, 2, 99, 85, 99, 30, 85, 85, 10}
	TriInsertionPlace(liste)
	if !listeEgal(liste, []int{-5, 2, 2, 10, 30, 30, 85, 85, 85, 99, 99}) {
		t.Fail()
	}
}

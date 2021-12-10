package algolp

import (
	"testing"
	"fmt"
)

func TestArbreBinaire(t *testing.T) {
	var b Arbre = GetEmptyArbre(1)

	if b.Root() != 1 || b.IsEmpty() {
		t.Fail()
	}
	if !b.Left().IsEmpty() || !b.Right().IsEmpty() {
		t.Fail()
	}
	var c Arbre = GetEmptyArbre(6)
	var d Arbre 
	d = MakeArbre(2, &b, &c)
	if d.Root() != 2 || d.left.Root() != 1 || d.right.IsEmpty() || !d.left.left.IsEmpty() {
		t.Fail()
	}
}

func TestArbreBinaireParcours(t *testing.T) {
	var b Arbre = GetEmptyArbre(1)
	var c Arbre = GetEmptyArbre(6)
	var d Arbre 
	d = MakeArbre(2, &b, &c)
	if d.Parcours() != "1 2 6" {
		t.Fail()
	}
}

func TestArbreBinaireRecherche(t *testing.T) {
	var b Arbre = GetEmptyArbre(1)
	var c Arbre = GetEmptyArbre(6)
	var d Arbre 
	d = MakeArbre(2, &b, &c)
	
	if d.Recherche(5) {
		t.Fail()
	}
	if !d.Recherche(1) || !d.Recherche(2) || !d.Recherche(6) {
		t.Fail()
	}
}

func TestArbreBinaireParcours2(t *testing.T) {
	var a Arbre = GetEmptyArbre(1)
	var b Arbre = GetEmptyArbre(2)
	var c = MakeArbre(3, &a, &b)

	var z Arbre = GetEmptyArbre(12)
	var d Arbre = MakeArbre(4, nil, &z)
	var e Arbre = MakeArbre(5, &d, nil)
	var f Arbre = MakeArbre(6, &c, &e)
	fmt.Println(f.Parcours())
	fmt.Println(f.Parcours2())
}
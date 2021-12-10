package algolp

import (
	"testing"
)

func TestArbreBinaire(t *testing.T) {
	var b Arbre = GetEmptyArbre(1)

	if b.Root() != 1 || b.IsEmpty() {
		t.Fail()
	}
	if !b.Left().IsEmpty() || !b.Right().IsEmpty() {
		t.Fail()
	}

	var d Arbre 
	d = Make(2, &b, nil)
	if d.Root() != 2 || d.left.Root() != 1 || !d.right.IsEmpty() || !d.left.left.IsEmpty() {
		t.Fail()
	}
	/*
	c := Make(3, &b, &d)
	d = Make(4, &b, &c)
	d.Parcours()
	*/

}

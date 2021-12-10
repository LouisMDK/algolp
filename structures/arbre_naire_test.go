package algolp

import "testing"

func TestArbreNaire(t *testing.T) {
	var a = GetEmptyArbreN(1)

	if a.Root() != 1 || len(a.nodes) != 0 {
		t.Fail()
	}
	b := GetEmptyArbreN(2)
	c := GetEmptyArbreN(3)
	d := GetEmptyArbreN(4)
	a = MakeArbreN(12, []*ArbreN{&b, &c, &d})

	if a.Root() != 12 || len(a.nodes) != 3 {
		t.Fail()
	}

	e := GetEmptyArbreN(16)
	f := GetEmptyArbreN(99)

	g := MakeArbreN(-1, []*ArbreN{&a, &e, &f})
	
	if g.Root() != -1 || len(g.nodes[0].nodes) != 3 || g.nodes[0].nodes[0].Root() != 2 {
		t.Fail()
	}
}

func TestArbreNaireParcours(t *testing.T) {
	var a = GetEmptyArbreN(1)
	b := GetEmptyArbreN(2)
	c := GetEmptyArbreN(3)
	d := GetEmptyArbreN(4)
	a = MakeArbreN(12, []*ArbreN{&b, &c, &d})
	e := GetEmptyArbreN(16)
	f := GetEmptyArbreN(99)
	g := MakeArbreN(-1, []*ArbreN{&a, &e, &f})

	if g.Parcours() != "-1 12 2 3 4 16 99" {
		t.Fail()
	}
}

func TestArbreNaireRecherche(t *testing.T) {
	var a = GetEmptyArbreN(1)
	b := GetEmptyArbreN(2)
	c := GetEmptyArbreN(3)
	d := GetEmptyArbreN(4)
	a = MakeArbreN(12, []*ArbreN{&b, &c, &d})
	e := GetEmptyArbreN(16)
	f := GetEmptyArbreN(99)
	g := MakeArbreN(-1, []*ArbreN{&a, &e, &f})

	if g.RechercheN(455) || !g.RechercheN(2) || !g.RechercheN(3) || !g.RechercheN(4) || !g.RechercheN(12) || !g.RechercheN(16) || !g.RechercheN(99) {
		t.Fail()
	}
}
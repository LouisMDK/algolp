package algolp

import "strconv"

type Arbre struct {
	val int
	left *Arbre
	right *Arbre
}

func (a Arbre) Root() int {
	return a.val
}

func (a Arbre) Left() *Arbre {
	return a.left
}

func (a Arbre) Right() *Arbre {
	return a.right
}

func MakeArbre(v int, a, b *Arbre) Arbre {
	return Arbre{val: v, left: a, right: b}
}

func GetEmptyArbre(v int) Arbre {
	return MakeArbre(v, nil, nil)
}

func (a *Arbre) IsEmpty() bool {
	return a == nil
}

func (a *Arbre) Recherche(v int) bool {
	if a.IsEmpty() {
		return false
	}
	if a.Root() == v {
		return true
	}
	l := a.left.Recherche(v)
	r := a.right.Recherche(v)
	return l || r
}

func (a *Arbre) Parcours() string {
	if !a.IsEmpty() {
		l := a.left.Parcours()
		if len(l) != 0 {
			l = l + " "
		}
		r := a.right.Parcours()
		if len(r) != 0 {
			r = " " + r
		}
		v := strconv.Itoa(a.val)
		d := l + v + r
		return d
	}
	return ""
}
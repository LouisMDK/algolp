package algolp

import "fmt"

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

func Make(v int, A *Arbre, B *Arbre) Arbre {
	return Arbre{val: v, left: A, right: B}
}

func GetEmptyArbre(v int) Arbre {
	return Make(v, nil, nil)
}

func (a *Arbre) IsEmpty() bool {
	return a == nil
}

func (a *Arbre) Parcours() {
	if !a.IsEmpty() {
		a.left.Parcours()
		fmt.Println(a.val)
		a.right.Parcours()
	}
}
package algolp

import "strconv"
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

func (a *Arbre) Taille() int {
	if a.IsEmpty() {
		return 0
	}
	return 1 + a.left.Taille() + a.right.Taille()
}

func (a *Arbre) Parcours2() string {
	if a.IsEmpty() {
		return ""
	}

	var vals [][]int = [][]int{}
	a.parcours2aux(0, &vals)
	fmt.Println(vals)
	var res string
	for i := 0; i < a.Taille(); i++ {
		for j := 0; j < len(vals); j++ {
			if vals[j][0] == i {
				if len(vals[j]) == 1 { // vide
					res = res + " "
				}else {
					res = res + strconv.Itoa(vals[j][1])
				}
			}
		}
		res = res + "\n"
	}
	return res
}

func (a *Arbre) parcours2aux(niv int, tab *[][]int) {
	c := []int{niv, a.val}
	fmt.Println(c)
	*tab = append(*tab, c)
	if !a.left.IsEmpty() {
		a.left.parcours2aux(niv + 1, tab)
	}else{
		*tab = append(*tab, []int{niv+1})
	}
	if !a.right.IsEmpty() {
		a.right.parcours2aux(niv + 1, tab)
	}else{
		*tab = append(*tab, []int{niv+1})
	}

	
}
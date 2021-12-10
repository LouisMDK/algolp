package algolp

import "strconv"

type ArbreN struct {
	val int
	nodes []*ArbreN
}

func (a ArbreN) Root() int {
	return a.val
}

func (a *ArbreN) IsEmpty() bool {
	return a == nil
}

func MakeArbreN(v int, fils []*ArbreN) ArbreN {
	return ArbreN{val: v, nodes: fils}
}

func GetEmptyArbreN(v int) ArbreN {
	return MakeArbreN(v, []*ArbreN{})
}

func (a *ArbreN) RechercheN(v int) bool {
	if a.IsEmpty() {
		return false
	}
	if a.Root()  == v {
		return true
	}
	
	for i := 0; i < len(a.nodes); i++ {
		if a.nodes[i].RechercheN(v) {
			return true
		}
	}
	return false
}

func (a *ArbreN) Parcours() (c string) {
	if !a.IsEmpty() {
		c += strconv.Itoa(a.val)
		for i := 0; i < len(a.nodes); i++ {
			text := a.nodes[i].Parcours()
			if len(text) != 0 {
				text = " " + text
			}
			c += text
		}
		return c
	}
	return ""
}
package algolp

import "fmt"

type Pile struct {
	element *pile_element
	cont int
}

type pile_element struct {
	v int
	next *pile_element
}

func GetEmptyPile() Pile{
	var pe pile_element = pile_element{next: nil}
	return Pile{element: &pe, cont: 0}
}

func (p *Pile) Push(val int) {
	var prochain *pile_element = p.element
	var pe pile_element = pile_element{v: val, next: prochain}
	p.element = &pe
	p.cont += 1
}

func (p Pile) IsEmpty() bool {
	return p.cont == 0
}

func (p *Pile) Pop() int {
	var val int = p.element.v
	p.element = p.element.next
	p.cont -= 1
	return val
}

func (p Pile) Affiche() string {
	var s string
	var l pile_element = *(p.element)
	for l.next != nil {
		s = fmt.Sprint(s, " ", l.v)
		l = *(l.next)
	}
	return s
}

func (p Pile) Count() int {
	return p.cont
}
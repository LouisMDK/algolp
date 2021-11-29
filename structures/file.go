package algolp

import "fmt"

type File struct {
	element *file_element
	cont int
	last *file_element
}

type file_element struct {
	v int
	next *file_element
	prev *file_element
}

func GetEmptyFile() File {
	var fe file_element = file_element{next: nil, prev: nil}
	return File{element: &fe, cont: 0, last: &fe}
}

func (f *File) Push(val int) {
	var new file_element = file_element{v: val, next: f.element, prev: nil}
	f.element.prev = &new
	f.element = &new
	f.cont++
	if f.cont == 1 {
		f.last = &new
	}
}

func (f File) IsEmpty() bool {
	return f.cont == 0
}

func (f *File) Pop() (res int) {
	if f.cont == 1 {
		res = f.last.v
		*f = GetEmptyFile()
		return res
	}
	res = f.last.v
	var new_final *file_element = f.last.prev
	f.last.prev.next = f.last.next
	f.last = new_final
	f.cont--
	return res
}
func (f File) Affiche() string {
	var s string
	var l file_element = *(f.element)
	for l.next != nil {
		s = fmt.Sprint(s, " ", l.v)
		l = *(l.next)
	}
	return s
}
func (f File) Count() int {
	return f.cont
}

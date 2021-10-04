package algolp

import (
	"testing"
)

func TestOccMax1(t *testing.T) {
	if _, occ := Occurencesmax(nil); occ != 0 {
		t.Fail()
	}
}

func TestOccMax2(t *testing.T) {
	if _, occ := Occurencesmax([]int{}); occ != 0 {
		t.Fail()
	}
}

func TestOccMax3(t *testing.T) {
	if _, occ := Occurencesmax([]int{1, 2, 3, 4, 5, 6, 7, 8}); occ != 1 {
		t.Fail()
	}
}

func TestOccMax4(t *testing.T) {
	if _, occ := Occurencesmax([]int{5, 5, 5, 5, 5}); occ != 5 {
		t.Fail()
	}
}

func TestOccMax5(t *testing.T) {
	if _, occ := Occurencesmax([]int{8, 1, 3, 4, 1, 6, 5, 2, 4, 5, 3, 5, 2, 6, 7, 8, 7}); occ != 2 {
		t.Fail()
	}
}

func TestValMax1(t *testing.T) {
	if n, _ := Occurencesmax([]int{5}); n != 5 {
		t.Fail()
	}
}

func TestVal2(t *testing.T) {
	if n, _ := Occurencesmax([]int{8, 1, 3, 4, 1, 6, 5, 2, 4, 5, 3, 5, 2, 6, 7, 8, 7}); n != 8 {
		t.Fail()
	}
}

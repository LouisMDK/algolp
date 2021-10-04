package algolp

import (
	"testing"
)

func TestOccurences1(t *testing.T) {
	if occurrences(5, nil) != 0 {
		t.Fail()
	}
}

func TestOccurences2(t *testing.T) {
	if occurrences(5, []int{}) != 0 {
		t.Fail()
	}
}

func TestOccurences3(t *testing.T) {
	if occurrences(5, []int{1, 2, 3, 4, 6, 7, 8}) != 0 {
		t.Fail()
	}
}

func TestOccurences4(t *testing.T) {
	if occurrences(5, []int{5, 5, 5, 5, 5}) != 5 {
		t.Fail()
	}
}

func TestOccurences5(t *testing.T) {
	if occurrences(5, []int{1, 5, 1, 1, 5, 5, 1, 5}) != 4 {
		t.Fail()
	}
}

func TestOccurences6(t *testing.T) {
	if occurrences(-5, []int{5, -5, 5, 5, -5}) != 2 {
		t.Fail()
	}
}

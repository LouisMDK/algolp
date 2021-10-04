package algolp

import (
	"testing"
)

func Test1Factorielle(t *testing.T) {
	if factorielle(0) != 1 {
		t.Fail()
	}
}

func Test2Factorielle(t *testing.T) {
	if factorielle(5) != 120 {
		t.Fail()
	}
}

func Test3Factorielle(t *testing.T) {
	if factorielle(10) != 3628800 {
		t.Fail()
	}
}

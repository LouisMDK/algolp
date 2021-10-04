package algolp

import (
	"testing"
)

func Test1(t *testing.T) {
	if EstPremier(0) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if EstPremier(1) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	if !EstPremier(2) {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	if !EstPremier(3) {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	if EstPremier(4) {
		t.Fail()
	}
}

func Test6(t *testing.T) {
	if !EstPremier(102647) {
		t.Fail()
	}
}

func Test7(t *testing.T) {
	if EstPremier(102647032) {
		t.Fail()
	}
}
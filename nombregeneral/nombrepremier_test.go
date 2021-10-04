package algolp

import (
	"testing"
)

func Test1NombrePremier(t *testing.T) {
	if EstPremier(0) {
		t.Fail()
	}
}

func Test2NombrePremier(t *testing.T) {
	if EstPremier(1) {
		t.Fail()
	}
}

func Test3NombrePremier(t *testing.T) {
	if !EstPremier(2) {
		t.Fail()
	}
}

func Test4NombrePremier(t *testing.T) {
	if !EstPremier(3) {
		t.Fail()
	}
}

func Test5NombrePremier(t *testing.T) {
	if EstPremier(4) {
		t.Fail()
	}
}

func Test6NombrePremier(t *testing.T) {
	if !EstPremier(102647) {
		t.Fail()
	}
}

func Test7NombrePremier(t *testing.T) {
	if EstPremier(102647032) {
		t.Fail()
	}
}
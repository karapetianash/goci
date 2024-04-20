package add

import "testing"

func TestAdd(t *testing.T) {
	a, b := 2, 3

	exp := 5

	res := add(a, b)

	if exp != res {
		t.Errorf("Expected %d, got %d instead.", exp, res)
	}
}

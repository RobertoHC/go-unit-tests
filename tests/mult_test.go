package tests

import(
	"testing"
	"go-unit-tests/functions"
)

func TestMult(t *testing.T) {
	total := functions.Mult(5, 5)
	if total != 24 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

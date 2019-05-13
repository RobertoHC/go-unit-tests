package tests

import(
	"testing"
	"go-unit-tests/functions"
)

func TestSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 3},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 3},
	}

	for _, table := range tables {
		total := functions.Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}
package collatz

import (
	"fmt"
	"testing"
)

// Collatz(2)
// 2, 1
//
// Collatz(3)
// 3, 16, 8, 4, 2, 1
func ExampleCollatz() {
	cnt, max, _ := Collatz(2)
	fmt.Printf("%d %d\n", cnt, max)

	cnt, max, _ = Collatz(3)
	fmt.Printf("%d %d\n", cnt, max)

	// Output:
	// 2 2
	// 8 16
}

func TestCollatz(t *testing.T) {
	ns := []uint64{0, 1, 2, 3}
	cnts := []uint64{0, 1, 2, 8}
	maxs := []uint64{0, 1, 2, 16}
	for i, n := range ns {
		got1, got2, _ := Collatz(n)
		want1, want2 := cnts[i], maxs[i]
		if got1 != want1 || got2 != want2 {
			t.Errorf("Collatz(%d) == %d, %d, want %d, %d",
				n, got1, got2, want1, want2)
		}
	}
}

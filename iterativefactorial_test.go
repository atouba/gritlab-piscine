package piscine

import "testing"

func TestIterativeFactorial(t *testing.T) {
	//   res := IterativeFactorial(4)
	//   if IterativeFactorial(4) != 4 {
	//     t.Errorf("Result is incorrect, got %d, want %d", 24, res)
	//   }
	for i := -5; i < 40; i++ {
		println(i, "!   = ", IterativeFactorial(i))
	}
}

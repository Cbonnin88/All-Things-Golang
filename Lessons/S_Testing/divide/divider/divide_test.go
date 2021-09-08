package divider

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	divide := Divide(3, 12)
	if divide != 4 {
		t.Error("Expected:", 4, "Got:", divide)
	}

}

func TestTableDivide(t *testing.T) {
	var tests = []struct {
		firstInput  int
		secondInput int
		answer      int
	}{
		{20, 100, 5},
		{6, 30, 5},
		{3, 12, 4},
		{4, 8, 2},
	}
	for _, test := range tests {
		if output := Divide(test.firstInput, test.secondInput); output != test.answer {
			t.Error("Expected:", test.answer, "Got:", output)
		}
	}
}

func ExampleDivide() {
	fmt.Println(Divide(25, 100))
	// Output:
	// 4
}

func BenchmarkDivide(b *testing.B) {
	for d := 0; d < b.N; d++ {
		Divide(3, 12)
	}
}

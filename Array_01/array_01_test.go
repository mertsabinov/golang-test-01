package array_01

import "testing"

func TestSum(t *testing.T) {

	numbers := []int{2, 2, 2, 2, 2}

	sum := Sum(numbers)
	want := 10

	if sum != want {
		t.Errorf("sum = %d  want = %d", sum, want)
	}

}

func BenchmarkSum(b *testing.B) {
	numbers := []int{2, 2, 2}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

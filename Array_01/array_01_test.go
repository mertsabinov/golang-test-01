package array_01

import (
	"reflect"
	"testing"
)

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

func TestSumsAll(t *testing.T) {

	sliceCheck := func(t *testing.T, sum []int, want []int) {
		if !reflect.DeepEqual(sum, want) {
			t.Errorf("sum = %v want = %v ", sum, want)
		}
	}

	t.Run("slice any size", func(t *testing.T) {
		sum := SumAll([]int{1, 1}, []int{2, 2})
		want := []int{2, 4}

		sliceCheck(t, sum, want)
	})

	t.Run("slice size zero", func(t *testing.T) {
		sum := SumAll([]int{}, []int{1, 2, 3, 4, 5})
		want := []int{0, 15}

		sliceCheck(t, sum, want)
	})
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{2, 2}, []int{1, 2, 3, 4, 5, 6})
	}
}

package struct_01

import "testing"

func TestArea(t *testing.T) {

	chackArea := func(got float64, want float64, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("got = %2f want = %2f", got, want)
		}
	}

	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 36.0},
		{name: "Circle", shape: Circle{Radius: 100}, hasArea: 628.00},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := Area(tt.shape)
			want := tt.hasArea
			chackArea(got, want, t)
		})
	}
}

func BenchmarkArea(b *testing.B) {
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 36.0},
		{name: "Circle", shape: Circle{Radius: 100}, hasArea: 628.00},
	}

	for _, tt := range areaTest {
		b.Run(tt.name, func(b *testing.B) {
			Area(tt.shape)
		})
	}
}

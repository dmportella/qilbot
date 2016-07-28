package edsm

import (
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	sothis := Coordinates{X: 0.0, Y: 0.0, Z: 0.0}
	bava := Coordinates{X: 1.0, Y: 1.0, Z: 1.0}
	expected := 1.732050807568877193176604123436845839023590087890625

	actual := calculateDistance(&sothis, &bava)

	if actual != expected {
		t.Errorf("calculateDistance(%v, %v): expected %g, actual %g", sothis, bava, expected, actual)
	}
}

func BenchmarkCalculateDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sothis := Coordinates{X: 0.0, Y: 0.0, Z: 0.0}
		bava := Coordinates{X: float64(n), Y: float64(n), Z: float64(n)}
		calculateDistance(&sothis, &bava)
	}
}

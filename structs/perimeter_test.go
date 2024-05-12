package perimeter

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		nameOfShape string
		shape       Shape
		hasArea     float64
	}{
		{nameOfShape: "Ractangle", shape: Ractangle{5.0, 5.0}, hasArea: 25.0},
		{nameOfShape: "Circle", shape: Circle{5}, hasArea: 78.53981633974483},
		{nameOfShape: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.nameOfShape, func(t *testing.T) {

			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.nameOfShape, got, tt.hasArea)
			}
		})
	}
}

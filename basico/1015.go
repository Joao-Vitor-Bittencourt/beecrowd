package basico

import (
	"fmt"
	"math"
)

func Bee1015() {
	var x1, x2, y1, y2 float64
	var distance float64

	fmt.Scanf("%g %g", &x1, &y1)
	fmt.Scanf("%g %g", &x2, &y2)

	distance = math.Sqrt(math.Pow(x2 - x1, 2.0) + math.Pow(y1 - y2, 2.0))

	fmt.Printf("%.4f", distance)
}
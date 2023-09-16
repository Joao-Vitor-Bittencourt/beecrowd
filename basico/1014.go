package basico

import "fmt"

func Bee1014() {
	var distance int
	var fuel float64

	fmt.Scanf("%d", &distance)
	fmt.Scanf("%f", &fuel)

	result := float64(distance) / fuel

	fmt.Printf("%.3f km/l\n", result)
}

package basico

import "fmt"

func Bee1011() {
	var radius float64

	fmt.Scanf("%f", &radius)

	result := (4.0 / 3.0) * 3.14159 * (radius * radius * radius)

	fmt.Printf("VOLUME = %.3f\n", result)
}

package basico

import (
	"fmt"
)

func Bee1006() {

	var a, b, c, result float64

	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)
	fmt.Scanf("%f", &c)

	result = ((a * 2) + (b * 3) + (c * 5)) / 10

	fmt.Printf("MEDIA = %.1f\n", result)
}

package basico

import "fmt"

func Bee1005() {
	var a, b, result float64

	fmt.Scanf("%g", &a)
	fmt.Scanf("%g", &b)

	result = ((3.5 * a) + (7.5 * b)) / (3.5 + 7.5)

	fmt.Printf("MEDIA = %.5f\n", result)
}

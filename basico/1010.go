package basico

import (
	"fmt"
)

func Bee1010() {
	var codeProduct1, codeProduct2 int
	var unitProduct1, unitProduct2, valueProduct1, valueProduct2 float64

	fmt.Scanf("%d %f %f", &codeProduct1, &unitProduct1, &valueProduct1)
	fmt.Scanf("%d %f %f", &codeProduct2, &unitProduct2, &valueProduct2)

	result := (unitProduct1 * valueProduct1) + (unitProduct2 * valueProduct2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", result)
}

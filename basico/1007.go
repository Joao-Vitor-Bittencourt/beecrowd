package basico

import "fmt"

func Bee1007() {
	var a, b, c, d, result int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &d)

	result = (a * b) - (c * d)

	fmt.Printf("DIFERENCA = %d\n", result)
}

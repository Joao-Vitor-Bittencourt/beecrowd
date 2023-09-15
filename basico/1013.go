package basico

import (
	"fmt"
	"math"
)

func Bee1013() {
	var a, b, c float64

	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)
	fmt.Scanf("%f", &c)

	result := (a + b + math.Abs(a-b)) / 2

	if result > c {
		fmt.Printf("%g eh o maior\n", result)
	} else {
		fmt.Printf("%g eh o maior\n", c)
	}
}

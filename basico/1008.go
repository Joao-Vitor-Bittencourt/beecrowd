package basico

import "fmt"

func Bee1008() {

	var number, hours int
	var value, result float64

	fmt.Scanf("%d", &number)
	fmt.Scanf("%d", &hours)
	fmt.Scanf("%f", &value)

	result = float64(hours) * value

	fmt.Printf("NUMBER = %d\n", number)
	fmt.Printf("SALARY = U$ %.2f\n", result)
}

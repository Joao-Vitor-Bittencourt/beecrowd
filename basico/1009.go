package basico

import "fmt"

func basico() {
	var sellersName string
	var fixedSalary, sales float64

	fmt.Scanf("%s", &sellersName)
	fmt.Scanf("%f", &fixedSalary)
	fmt.Scanf("%f", &sales)

	result := fixedSalary + ((sales * 15) / 100)

	fmt.Printf("TOTAL = R$ %.2f\n", result)
}

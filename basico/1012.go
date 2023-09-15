package basico

import (
	"fmt"
)

func Bee1012() {

	var a, b, c float64

	fmt.Scanf("%f %f %f", &a, &b, &c)

	pi := 3.14159

	triangulo := a * c / 2
	circulo := (c * c) * pi
	trapezio := (a + b) * c / 2
	quadrado := b * b
	retangulo := a * b

	fmt.Printf("TRIANGULO: %.3f\n", triangulo)
	fmt.Printf("CIRCULO: %.3f\n", circulo)
	fmt.Printf("TRAPEZIO: %.3f\n", trapezio)
	fmt.Printf("QUADRADO: %.3f\n", quadrado)
	fmt.Printf("RETANGULO: %.3f\n", retangulo)
}

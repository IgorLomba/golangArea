package golangArea

import "math"

/* pacote area, pasta area */

const Pi = math.Pi

/* Circulo ... */
func Circulo(raio float64) (area float64) {
	area = math.Pow(raio, 2) * Pi
	return
}

/* Retangulo ... */
func Rectangulo(base, altura float64) (area float64) {
	area = base * altura
	return
}

/* Triangulo não é visível externamente, as outras são */
func triangulo(base, altura float64) (area float64) {
	area = (base * altura) / 2
	return
}

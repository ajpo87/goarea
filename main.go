package goarea

import "math"

//PI proporcao numerica constante
const Pi = 3.1415

//Funcao area circulo
func Circulo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Area rectangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//nao e funcao visivel
func _TringuloEq(base, altura float64) float64 {
	return (base * altura) / 2
}

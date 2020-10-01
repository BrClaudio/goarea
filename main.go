package goarea

import "math"

//Pi comentario obrigatorio
const Pi = 3.1416

//Circulo fgfgfgfde novo obrigartorio pq publico
func Circulo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect kfgjfkgjfkgjfkfkg ffkgj
func Rect(base, altura float64) float64 {
	return base * altura
}

//Triangulo fkjgjfkgj fkgj fkgjfk j
func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}

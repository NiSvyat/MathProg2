package main

import (
	"fmt"
)

type Point struct {
	X, Y float64
}

// Функция для вычисления значения многочлена Лагранжа в точке x
func lagrangeInterpolation(points []Point, x float64) float64 {
	var result float64
	for i := 0; i < len(points); i++ {
		term := points[i].Y // Начинаем с y_i
		for j := 0; j < len(points); j++ {
			if i != j { // Пропускаем текущий i
				term *= (x - points[j].X) / (points[i].X - points[j].X)
			}
		}
		result += term
	}
	return result
}

func main() {
	// Исходные данные из таблицы
	points := []Point{
		{-2, 25},
		{1, -8},
		{2, -15},
		{4, -23},
	}

	// Точка, в которой нужно вычислить значение многочлена
	x := 3.0

	// Вычисляем значение многочлена Лагранжа в точке x
	interpolatedY := lagrangeInterpolation(points, x)

	// Выводим результат
	fmt.Printf("Значение многочлена Лагранжа в точке x = %.2f: %.4f\n", x, interpolatedY)
}

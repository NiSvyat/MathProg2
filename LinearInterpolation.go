package main

import (
	"fmt"
	"math"
)

func linearInterpolation(x0, y0, x1, y1, x float64) float64 {
	return y0 + (y1-y0)*(x-x0)/(x1-x0)
}

func main() {
	// Исходные данные
	x := []float64{4, 9, 16, 25}
	y := []float64{2, 3, 4, 5}

	// Точка, в которой нужно найти значение функции
	targetX := 11.0

	// Находим ближайшие точки для интерполяции
	var x0, y0, x1, y1 float64
	for i := 0; i < len(x)-1; i++ {
		if x[i] <= targetX && x[i+1] >= targetX {
			x0, y0 = x[i], y[i]
			x1, y1 = x[i+1], y[i+1]
			break
		}
	}

	// Линейная интерполяция
	interpolatedY := linearInterpolation(x0, y0, x1, y1, targetX)

	// Оценка погрешности
	exactY := math.Sqrt(targetX)
	error := math.Abs(interpolatedY - exactY)

	fmt.Printf("Интерполированное значение y(%.2f) = %.4f\n", targetX, interpolatedY)
	fmt.Printf("Точное значение y(%.2f) = %.4f\n", targetX, exactY)
	fmt.Printf("Погрешность: %.4f\n", error)
}

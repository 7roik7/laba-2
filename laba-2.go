package main

import (
	"fmt"
	"unicode/utf8"
)

// 1. Чётное / нечётное
func evenOdd(n int) string {
	if n%2 == 0 {
		return "чётное"
	}
	return "нечётное"
}

// 2. Positive / Negative / Zero
func sign(n int) string {
	switch {
	case n > 0:
		return "Positive"
	case n < 0:
		return "Negative"
	}
	return "Zero"
}

// 4. Длина строки
func length(s string) int { return utf8.RuneCountInString(s) }

// 5. Структура Rectangle + метод площади
type Rectangle struct{ W, H float64 }

func (r Rectangle) Area() float64 { return r.W * r.H }

// 6. Среднее двух целых
func avg(a, b int) float64 { return float64(a+b) / 2 }

func main() {
	// 1
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	fmt.Println("Число", evenOdd(n))

	// 2
	fmt.Println(sign(n))

	// 3
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 4
	fmt.Println("Длина:", length("Привет"))

	// 5
	r := Rectangle{5, 3}
	fmt.Println("Площадь:", r.Area())

	// 6
	fmt.Println("Среднее:", avg(4, 9))
}

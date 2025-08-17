package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тест функции generateRandomElements
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  []int
	}{
		{"Размер 0", 0, []int{}},
		{"Размер 1", 1, []int{}},
		{"Размер 10", 10, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.input)
			assert.Equalf(t, tt.want, len(got), "Длина результирующего массива отличается от ожидания")
		})
	}
}

// Тест функции maximum
func TestMaximum(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Пустой слайс", args{[]int{}}, 0},
		{"Один элемент", args{[]int{1}}, 1},
		{"Малый набор", args{[]int{1, 2, 3}}, 3},
		{"Средний набор", args{[]int{1, 2, 3, 4, 5}}, 5},
		{"Большой набор", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, 10},
		{"Максимальное число", args{[]int{1, 2, 3, 1000}}, 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.args.data)
			assert.Equalf(t, tt.want, got, "Ошибка в поиске максимума: want=%v, got=%v", tt.want, got)
		})
	}
}

// Тест функции maxChunks
func TestMaxChunks(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Пустой слайс", args{[]int{}}, 0},
		{"Один элемент", args{[]int{1}}, 1},
		{"Малый набор", args{[]int{1, 2, 3}}, 3},
		{"Средний набор", args{[]int{1, 2, 3, 4, 5}}, 5},
		{"Большой набор", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, 10},
		{"Граничные числа", args{[]int{1, 1000, 500, 10000}}, 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.args.data)
			assert.Equalf(t, tt.want, got, "Ошибка в поиске максимума в сегментах: want=%v, got=%v", tt.want, got)
		})
	}
}

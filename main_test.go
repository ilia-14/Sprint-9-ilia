package main

import (
	"testing"
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
			if len(got) != tt.input {
				t.Errorf("generateRandomElements(%d) length mismatch: want=%d, got=%d", tt.input, tt.input, len(got))
			}
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
		{"Пустой слайс", args{[]int{}}, 0},                   // Паника
		{"Один элемент", args{[]int{1}}, 1},                  // Максимум очевиден
		{"Несколько элементов", args{[]int{-1, 0, 3, 2}}, 3}, // Наибольшее число
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.args.data); got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
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
		{"Пустой слайс", args{[]int{}}, 0},                   // Паника
		{"Один элемент", args{[]int{1}}, 1},                  // Максимум единстенной чисто
		{"Несколько элементов", args{[]int{-1, 0, 3, 2}}, 3}, // Наибольшее число
		{"Большое количество элементов", args{generateRandomElements(10)}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.args.data); got != tt.want {
				t.Errorf("maxChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

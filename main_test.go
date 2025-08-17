package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тест функции generateRandomElements
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Размер 0", 0, 0},
		{"Размер 1", 1, 1},
		{"Размер 10", 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.input)
			assert.Len(t, got, tt.want, "Длина результирующего массива отличается от ожидания")
		})
	}
}

// Тест функции maximum
func TestMaximum(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Пустой слайс", args{[]int{}}, 0, true},
		{"Один элемент", args{[]int{1}}, 1, false},
		{"Малый набор", args{[]int{1, 2, 3}}, 3, false},
		{"Средний набор", args{[]int{1, 2, 3, 4, 5}}, 5, false},
		{"Большой набор", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, 10, false},
		{"Максимальное число", args{[]int{1, 2, 3, 1000}}, 1000, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := maximum(tt.args.data)
			if tt.wantErr {
				require.Error(t, err, "Ожидается ошибка")
			} else {
				require.NoError(t, err, "Ошибка не ожидалась")
				assert.Equal(t, tt.want, got, "Полученное значение не совпадает с ожидаемым")
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
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Пустой слайс", args{[]int{}}, 0, true},
		{"Один элемент", args{[]int{1}}, 1, false},
		{"Малый набор", args{[]int{1, 2, 3}}, 3, false},
		{"Средний набор", args{[]int{1, 2, 3, 4, 5}}, 5, false},
		{"Большой набор", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, 10, false},
		{"Граничные числа", args{[]int{1, 1000, 500, 10000}}, 10000, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := maxChunks(tt.args.data)
			if tt.wantErr {
				require.Error(t, err, "Ожидается ошибка")
			} else {
				require.NoError(t, err, "Ошибка не ожидалась")
				assert.Equal(t, tt.want, got, "Полученное значение не совпадает с ожидаемым")
			}
		})
	}
}

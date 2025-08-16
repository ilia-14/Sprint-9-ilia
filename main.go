package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	rand.Seed(time.Now().Unix()) // Засеваем генератор случайных чисел для уникальности
	result := make([]int, size)
	for i := range size {
		result[i] = rand.Int() // Генерируем случайное целое число
	}
	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		panic("Empty slice provided.")
	}
	maxValue := data[0]
	for _, num := range data {
		if num > maxValue {
			maxValue = num
		}
	}
	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	size := len(data)
	if size == 0 {
		panic("Empty slice provided.")
	}

	chunkSize := size / CHUNKS
	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		startIndex := i * chunkSize
		endIndex := startIndex + chunkSize
		if i == CHUNKS-1 {
			endIndex = size // Последний кусок может быть меньше стандартного
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			localMax := data[start]
			for j := start + 1; j < end; j++ {
				if data[j] > localMax {
					localMax = data[j]
				}
			}
			maxValues[i] = localMax
		}(startIndex, endIndex)
	}

	wg.Wait()

	globalMax := maxValues[0]
	for _, value := range maxValues {
		if value > globalMax {
			globalMax = value
		}
	}
	return globalMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	startTime := time.Now()
	maxSingleThread := maximum(numbers)
	elapsed := time.Since(startTime).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxSingleThread, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startTime = time.Now()
	maxMultiThread := maxChunks(numbers)
	elapsed = time.Since(startTime).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxMultiThread, elapsed)
}

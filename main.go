package main

import (
	"errors"
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
	result := make([]int, size)
	for i := range size {
		result[i] = rand.Intn(10000) // Генерируем случайное целое число
	}
	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("пустой срез предоставлен")
	}
	maxValue := data[0]
	for _, num := range data {
		if num > maxValue {
			maxValue = num
		}
	}
	return maxValue, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {
	size := len(data)
	if size == 0 {
		return 0, errors.New("пустой срез предоставлен")
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

		chunkData := data[startIndex:endIndex]

		wg.Add(1)
		go func(chunk []int, i int) {
			defer wg.Done()
			maxValues[i], _ = maximum(chunk)
		}(chunkData, i)
	}

	wg.Wait()

	globalMax, _ := maximum(maxValues)
	return globalMax, nil
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	startTime := time.Now()
	maxSingleThread, _ := maximum(numbers)
	elapsed := time.Since(startTime).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxSingleThread, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startTime = time.Now()
	maxMultiThread, _ := maxChunks(numbers)
	elapsed = time.Since(startTime).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxMultiThread, elapsed)
}

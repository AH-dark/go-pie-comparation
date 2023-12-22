package main

import (
	"fmt"
	"math"
	"os"
	"sync"
)

func compute(start, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	var result float64
	for i := start; i < end; i++ {
		num := math.Sqrt(float64(i)) * math.Sin(float64(i)) * math.Cos(float64(i))
		result += num
	}
}

func main() {
	const numWorkers = 4
	const numElements = 25000000
	wg := sync.WaitGroup{}

	// 并发数学计算
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go compute(i*numElements/numWorkers, (i+1)*numElements/numWorkers, &wg)
	}

	// 内存操作
	data := make(map[int]int)
	for i := 0; i < 10000000; i++ {
		data[i] = i ^ 0xff00
	}

	// 文件I/O
	file, err := os.CreateTemp("", "go-pie-test-*.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Test data")

	wg.Wait()
	fmt.Println("Test completed")
}

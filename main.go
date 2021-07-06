package main

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/loov/hrtime"
)

func func_call(n int) {
	mock_value := int64(0)
	for i := 0; i < n; i++ {
		mock_value += int64(i)
	}
}

func test1() {
	n := 1000000 // 10**6

	debug.SetGCPercent(-1)

	// Test func call
	b1 := hrtime.NewBenchmark(100)
	for b1.Next() {
		func_call(n)
	}
	_, err := fmt.Println(b1.Histogram(10))
	if err == nil {
	} else {
		// Если раскомментировать строчку снизу, то перфоманс второго цикла увеличится вдвое? Почему?
		// OS: Windows 10,
		// Выполнение команды
		// 		$wmic cpu get caption, deviceid, name, numberofcores, maxclockspeed, status
		// дает вывод
		// 		> AMD64 Family 23 Model 113 Stepping 0 CPU0 3793 AMD Ryzen 9 3900X 12-Core Processor 12 OK
		// fmt.Println("Error is occured")
	}

	b2 := hrtime.NewBenchmark(500)
	for b2.Next() {
		// this is the body of func_call
		mock_value := int64(0)
		for i := 0; i < n; i++ {
			mock_value += int64(i)
		}
	}

	fmt.Println(b2.Histogram(10))
}

func test2() {
	n := 1000000 // 10**6
	bench_count := 100
	debug.SetGCPercent(-1)

	// Test func call
	start_time := time.Now()
	for i := 0; i < bench_count; i++ {
		func_call(n)
	}
	elapsed := time.Since(start_time)
	fmt.Println("Elapsed average", elapsed/time.Duration(bench_count))

	// fmt.Println("Testing inline approach")
	start_time = time.Now()
	for j := 0; j < bench_count; j++ {
		mock_value := int64(0)
		for i := 0; i < n; i++ {
			mock_value += int64(i)
		}
	}
	elapsed = time.Since(start_time)
	fmt.Println("Elapsed average", elapsed/time.Duration(bench_count))
}

func main() {
	test1()
	// test2()
}

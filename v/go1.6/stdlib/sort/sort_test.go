package main

import (
	"log"
	"math/rand"
	"runtime"
	"sort"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	log.Println(runtime.Version())
}

func BenchmarkSort_1(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1)
		b.StartTimer()
		sort.Ints(data)
		b.StopTimer()
	}
}

// func BenchmarkSort_10(b *testing.B) {
// 	b.StopTimer()
// 	for i := 0; i < b.N; i++ {
// 		data := make([]int, 10)
// 		for i := 0; i < len(data); i++ {
// 			data[i] = rand.Intn(len(data))
// 		}
// 		b.StartTimer()
// 		sort.Ints(data)
// 		b.StopTimer()
// 	}
// }
// func BenchmarkSort_100(b *testing.B) {
// 	b.StopTimer()
// 	for i := 0; i < b.N; i++ {
// 		data := make([]int, 100)
// 		for i := 0; i < len(data); i++ {
// 			data[i] = rand.Intn(len(data))
// 		}
// 		b.StartTimer()
// 		sort.Ints(data)
// 		b.StopTimer()
// 	}
// }
// func BenchmarkSort_1000(b *testing.B) {
// 	b.StopTimer()
// 	for i := 0; i < b.N; i++ {
// 		data := make([]int, 1000)
// 		for i := 0; i < len(data); i++ {
// 			data[i] = rand.Intn(len(data))
// 		}
// 		b.StartTimer()
// 		sort.Ints(data)
// 		b.StopTimer()
// 	}
// }

// func BenchmarkSort_10000(b *testing.B) {
// 	b.StopTimer()
// 	for i := 0; i < b.N; i++ {
// 		data := make([]int, 10000)
// 		for i := 0; i < len(data); i++ {
// 			data[i] = rand.Intn(len(data))
// 		}
// 		b.StartTimer()
// 		sort.Ints(data)
// 		b.StopTimer()
// 	}
// }

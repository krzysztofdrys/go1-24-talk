package maps

import (
	"fmt"
	"testing"
)

func BenchmarkMapAppend(b *testing.B) {
	testSize := []int64{10, 100, 1000, 10000}

	for _, size := range testSize {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m := make(map[int64]struct{}, 0)
				for j := int64(0); j < size; j++ {
					m[j] = struct{}{}
				}
			}
		})
	}
}

func BenchmarkMapInsert(b *testing.B) {
	testSize := []int64{10, 100, 1000, 10000}

	for _, size := range testSize {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m := make(map[int64]struct{}, size)
				for j := int64(0); j < size; j++ {
					m[j] = struct{}{}
				}
			}
		})
	}
}

func BenchmarkMapGetHit(b *testing.B) {
	testSize := []int64{10, 100, 1000, 10000}
	var sink int

	for _, size := range testSize {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			b.StopTimer()
			m := make(map[int64]struct{}, size)
			for j := int64(0); j < size; j++ {
				m[int64(j)] = struct{}{}
			}
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				for j := int64(0); j < size; j++ {
					_, ok := m[j]
					if ok {
						sink++
					}
				}
			}
		})
	}
}

func BenchmarkMapGetMiss(b *testing.B) {
	testSize := []int64{10, 100, 1000, 10000}
	var sink int

	for _, size := range testSize {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			b.StopTimer()
			m := make(map[int64]struct{}, size)
			for j := int64(0); j < size; j++ {
				m[j] = struct{}{}
			}
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				for j := int64(0); j < size; j++ {
					_, ok := m[j+size]
					if ok {
						sink++
					}
				}
			}
		})
	}
}

func BenchmarkMapDelete(b *testing.B) {
	testSize := []int64{10, 100, 1000, 10000}

	for _, size := range testSize {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			b.StopTimer()
			m := make(map[int64]struct{}, size)
			for j := int64(0); j < size; j++ {
				m[j] = struct{}{}
			}
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				for j := int64(0); j < size; j++ {
					delete(m, j)
				}
			}
		})
	}
}

func BenchmarkMapCreate(b *testing.B) {
	testSize := []int64{10, 100, 1000, 10000}
	sink := 0

	for _, size := range testSize {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m := make(map[int64]struct{}, size)
				sink += len(m)
			}
		})
	}
}

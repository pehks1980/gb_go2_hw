package main

import (
	"sync"
	"testing"
)

// структура и методы чт/зп RW mutex
type RWSet struct {
	sync.RWMutex
	mm map[int]struct{}
}
// конструктор множества RW mutex
func NewRWSet() *RWSet {
	return &RWSet{
		mm: map[int]struct{}{},
	}
}
// Зп множества RW mutex
func(s *RWSet) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}
// Чт множества RW mutex
func(s *RWSet) Has(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}

// структура и методы чт/зп mutex
type Set struct {
	sync.Mutex
	mm map[int]struct{}
}
// конструктор множества mutex
func NewSet() *Set {
	return &Set{
		mm: map[int]struct{}{},
	}
}
// Зп множества mutex
func(s *Set) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}
// Чт множества mutex
func(s *Set) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

func BenchmarkSetAdd(b *testing.B) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}

/*
Протестируйте производительность множества действительных чисел, безопасность которого обеспечивается
sync.Mutex и sync.RWMutex для разных вариантов использования:
10% запись, 90% чтение;
50% запись, 50% чтение;
90% запись, 10% чтение
 */

func BenchmarkSetHas(b *testing.B) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkRWSetAdd(b *testing.B) {
	var set = NewRWSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}

func BenchmarkRWSetHas(b *testing.B) {
	var set = NewRWSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

//go test -bench=. main_test.go

//10 wr 90 rd RW set
func BenchmarkRWSet10WR90RD(b *testing.B) {
	var (
		setR = NewRWSet()
		setWR = NewRWSet()
	)
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i:=0;i<10;i++{
					setWR.Add(i)
				}
				for i:=0;i<90;i++{
					setR.Has(i%10)
				}

			}
		})
	})
}

//10 wr 90 rd set
func BenchmarkSet10WR90RD(b *testing.B) {
	var (
		setR = NewSet()
		setWR = NewSet()
	)
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i:=0;i<10;i++{
					setWR.Add(i)
				}
				for i:=0;i<90;i++{
					setR.Has(i%10)
				}

			}
		})
	})
}

//50 wr 50 rd RW set
func BenchmarkRWSet50WR50RD(b *testing.B) {
	var (
		setR = NewRWSet()
		setWR = NewRWSet()
	)
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i:=0;i<50;i++{
					setWR.Add(i)
				}
				for i:=0;i<50;i++{
					setR.Has(i)
				}

			}
		})
	})
}

//50 wr 50 rd set
func BenchmarkSet50WR50RD(b *testing.B) {
	var (
		setR = NewSet()
		setWR = NewSet()
	)
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i:=0;i<50;i++{
					setWR.Add(i)
				}
				for i:=0;i<50;i++{
					setR.Has(i)
				}

			}
		})
	})
}

//10 rd 90 wr RW set
func BenchmarkRWSet10RD90WR(b *testing.B) {
	var (
		setR = NewRWSet()
		setWR = NewRWSet()
	)
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i:=0;i<10;i++{
					setR.Has(i)
				}
				for i:=0;i<90;i++{
					setWR.Add(i)
				}

			}
		})
	})
}

//10 rd 90 wr set
func BenchmarkSet10RD90WR(b *testing.B) {
	var (
		setR = NewSet()
		setWR = NewSet()
	)
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i:=0;i<10;i++{
					setR.Has(i)
				}
				for i:=0;i<90;i++{
					setWR.Add(i)
				}

			}
		})
	})
}
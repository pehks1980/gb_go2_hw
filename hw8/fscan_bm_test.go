package main_test

import (
	"runtime"
	"strconv"
	"testing"
	mn "github.com/pehks1980/gb_go2_hw/hw8"
)
// go test -bench=.

func BenchmarkScanDir1CPU(b *testing.B) {
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mn.ScanDir("/home/user/go")
			}
		})
	})
}

func BenchmarkScanDir1CPU_1(b *testing.B) {
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mn.ScanDir("/home/user")
			}
		})
	})
}

func BenchmarkScanDir4CPU(b *testing.B) {
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(4)), func(b *testing.B) {
		b.SetParallelism(10000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mn.ScanDir("/home/user/go")
			}
		})
	})
}

func BenchmarkScanDir4CPU_1(b *testing.B) {
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(4)), func(b *testing.B) {
		b.SetParallelism(10000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mn.ScanDir("/home/user")
			}
		})
	})
}
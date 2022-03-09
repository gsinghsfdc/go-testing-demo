package fib

import (
	"strconv"
	"testing"
)

var (
	basicCases = map[int]int{
		0:  0,
		1:  1,
		2:  1,
		3:  2,
		4:  3,
		5:  5,
		10: 55,
	}
	benchCases = []int{1, 10, 40}
)

func TestFibRecursive(t *testing.T) {
	for i, v := range basicCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if f := FibRecursive(i); f != v {
				t.Fatalf("FibRecursive(%v) == %v : expected %v", i, f, v)
			}
		})
	}
}

func TestFibCache(t *testing.T) {
	for i, v := range basicCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if f := FibCache(i); f != v {
				t.Fatalf("FibCache(%v) == %v : expected %v", i, f, v)
			}
		})
	}
}

func BenchmarkFibRecursive(b *testing.B) {
	for _, t := range benchCases {
		tc := func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				FibRecursive(t)
			}
		}
		b.Run(strconv.Itoa(t), tc)
	}
}

func BenchmarkFibCache(b *testing.B) {
	for _, t := range benchCases {
		tc := func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				FibCache(t)
			}
		}
		b.Run(strconv.Itoa(t), tc)
	}
}

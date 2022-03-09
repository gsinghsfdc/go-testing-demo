package fib

import (
	"strconv"
	"testing"
)

func TestFibRecursive(t *testing.T) {
	tests := map[int]int{
		0:  0,
		1:  1,
		2:  1,
		3:  2,
		4:  3,
		5:  5,
		10: 55,
	}
	for i, v := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if f := FibRecursive(i); f != v {
				t.Fatalf("FibRecursive(%v) == %v : expected %v", i, f, v)
			}
		})
	}
}

func TestFibCache(t *testing.T) {
	tests := map[int]int{
		0:  0,
		1:  1,
		2:  1,
		3:  2,
		4:  3,
		5:  5,
		10: 55,
	}
	for i, v := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if f := FibCache(i); f != v {
				t.Fatalf("FibCache(%v) == %v : expected %v", i, f, v)
			}
		})
	}
}

func BenchmarkFibRecursive(b *testing.B) {
	ts := []int{1, 10, 40}
	for _, t := range ts {
		tc := func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				FibRecursive(t)
			}
		}
		b.Run(strconv.Itoa(t), tc)
	}
}

func BenchmarkFibCache(b *testing.B) {
	ts := []int{1, 10, 40}
	for _, t := range ts {
		tc := func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				FibCache(t)
			}
		}
		b.Run(strconv.Itoa(t), tc)
	}
}

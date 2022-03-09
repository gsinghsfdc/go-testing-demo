package fib

func FibRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return FibRecursive(n-1) + FibRecursive(n-2)
}

var cache = map[int]int{0: 0, 1: 1}

func FibCache(n int) int {
	if f, ok := cache[n]; ok {
		return f
	}

	cache[n] = FibCache(n-1) + FibCache(n-2)

	return cache[n]
}

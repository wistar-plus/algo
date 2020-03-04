package recursive

func Fibonacci(n uint) uint {
	if n == 0 || n == 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func ClimbingStairs(n uint) uint {
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	return ClimbingStairs(n-1) + ClimbingStairs(n-2)
}

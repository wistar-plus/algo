package recursive

import (
	"fmt"
	"testing"
)

func Test_Fibonacci(t *testing.T) {
	for i := uint(0); i < 10; i++ {
		fmt.Printf("%v ", Fibonacci(i))
	}
	fmt.Println()
}

func Test_ClimbingStairs(t *testing.T) {
	for i := uint(0); i < 10; i++ {
		fmt.Printf("%v ", ClimbingStairs(i))
	}
	fmt.Println()
}

package main

import (
	"fmt"
	"log"
)

const maxUint64 = 1<<64 - 2
var F = make(map[uint64]uint64)

// Complexity: O(logn) because it divides n to half for every calculation
func fib(n uint64) uint64 {
	var k uint64
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		F[n] = 1
		goto RETURN
	}
	if F[n] > 0 {
		goto RETURN
	}
	if n%2 == 0 {
		k = n/2
		F[n] = (2*fib(k-1) + fib(k))*fib(k)
	} else {
		k = (n+1)/2
		F[n] = fib(k)*fib(k) + fib(k-1)*fib(k-1)
	}
	RETURN:
	return F[n]
}

func mod(n, p uint64) uint64 {
	if n > maxUint64 {
		log.Fatal("maximum allowed exceed")
	}
	result := fib(n)
	fmt.Println(fmt.Sprintf("fibonacci=%v", result))
	return result%p
}

func main() {
	fmt.Println(mod(6, 3))
	fmt.Println(mod(100, 2))
	fmt.Println(mod(maxUint64, 3))
}

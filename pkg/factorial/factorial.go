package factorial

import (
	"math"
	"sync"
)

func Factorial(n int, wg *sync.WaitGroup, result *int) {
	defer wg.Done()
	fact := 1

	for i := 2; i <= n; i++ {
		if fact > math.MaxInt32/i {
			*result = 0
			return
		}
		fact *= i
	}

	*result = fact
}

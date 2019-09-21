package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	// Parse input value from command line args.  Do simple validation.
	n := 10
	if len(os.Args) > 1 {
		parsed, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil || parsed < 1 {
			panic("Input target must be a positive integer")
		}
		n = int(parsed)
	}

	sum := big.NewRat(0, 1)

	for i := 1; i <= n; i++ {
		sum = sum.Add(sum, big.NewRat(1, int64(i)))
	}

	fmt.Printf("Frog(%d) = %s == %s\n", n, sum.String(), sum.FloatString(10))
}

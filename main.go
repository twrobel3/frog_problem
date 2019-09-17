package main

import (
	"encoding/csv"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
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

	// Likelihoods is a slice that's built up iteratively, as the expected value
	// on lillypad n is dependent on the expected value for all lillypads from
	// 1 to n-1.
	likelihoods := make([]*big.Rat, n+1)
	likelihoods[0] = big.NewRat(0, 1)

	w := csv.NewWriter(os.Stdout)

	for i := 1; i <= n; i++ {
		expected := big.NewRat(0, 1)
		for j := 0; j < i; j++ {
			value := big.NewRat(0, 1)
			value.Add(big.NewRat(1, 1), likelihoods[j])
			value.Mul(value, big.NewRat(1, int64(i)))

			expected.Add(expected, value)
		}
		likelihoods[i] = expected

		numDenom := strings.Split(expected.String(), "/")
		f, _ := expected.Float64()
		w.Write([]string{fmt.Sprintf("%d", i), numDenom[0], numDenom[1], fmt.Sprintf("%f", f)})
	}

	w.Flush()
}

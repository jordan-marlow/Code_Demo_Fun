package main

import (
	"fmt"
	"math/big"
	"time"
)

func init_primes(max_num int64) []int64 {
	primes := make([]bool, max_num+1)

	//initialize primes boolean array to true
	for i := int64(2); i <= max_num; i++ {
		primes[i] = true
	}

	//this is the sieve method.  If the element of primes is true, then every multiple of that prime henceforth is set to false.
	for p := int64(2); p*p <= max_num; p++ {
		if primes[p] {
			for i := p * p; i <= max_num; i += p {
				primes[i] = false
			}
		}
	}

	var primeNumbers []int64
	for p := int64(2); p <= max_num; p++ {
		if primes[p] {
			primeNumbers = append(primeNumbers, p)
		}
	}
	return primeNumbers
}

func lucas_lehmer_test(prime int64) bool {
	if prime == 2 {
		return true
	}
	s := big.NewInt(4)

	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)
	big_prime := big.NewInt(prime)
	M := new(big.Int).Sub(new(big.Int).Exp(two, big_prime, nil), one)

	for i := int64(0); i < prime-1; i++ {
		s.Mul(s, s).Sub(s, two).Mod(s, M)
		if s.Cmp(zero) == 0 {
			return true
		}
	}
	return false
}

func main() {
	var maxNum int64
	fmt.Println("Enter the maximum number to check for Mersenne Primes.  [5000 is a good number that takes about 5s to complete.]: ")
	_, err := fmt.Scanln(&maxNum)
	if err != nil {
		fmt.Println("Invalid input.  Re-run and enter a valid integer.")
	}
	primes := init_primes(maxNum)
	now := time.Now()
	for _, p := range primes {
		if lucas_lehmer_test(p) {
			fmt.Printf("2^%d - 1 is a Mersenne Prime and took %v to find!\n", p, time.Since(now))
		}
	}
	fmt.Println("Press enter to exit...")
	fmt.Scanln()
}

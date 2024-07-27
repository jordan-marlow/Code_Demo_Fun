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

func lucas_lehmer_test(prime int64, resultChannel chan<- string) {
	start_time := time.Now()
	if prime == 2 {
		resultChannel <- fmt.Sprintf("2^%d - 1 is a Mersenne Prime and took %v to find!", prime, time.Since(start_time))
		return
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
			resultChannel <- fmt.Sprintf("2^%d - 1 is a Mersenne Prime and took %v to find!", prime, time.Since(start_time))
			return
		}
	}
	resultChannel <- ""
}

func main() {
	var maxNum int64
	fmt.Println("Enter the maximum number to check for Mersenne Primes.  [5000 is a good number that takes about 5s to complete.]: ")
	_, err := fmt.Scanln(&maxNum)
	if err != nil {
		fmt.Println("Invalid input.  Re-run and enter a valid integer.")
	}
	primes := init_primes(maxNum)
	resultChannel := make(chan string)
	for _, p := range primes {
		go lucas_lehmer_test(p, resultChannel)
	}

	for i := 0; i < len(primes); i++ {
		result := <-resultChannel
		if result != "" {
			fmt.Println(result)
		}
	}
	close(resultChannel)
	fmt.Println("Press enter to exit...")
	fmt.Scanln()
}

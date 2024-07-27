package main

import (
	"fmt"
	"math/big"
)

// The chudnovsky algorithm can be found here [https://en.wikipedia.org/wiki/Chudnovsky_algorithm]

// Recursive Factorial Function
func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}

func power(x *big.Int, y int64) *big.Int {
	n := big.NewInt(1)
	if y == 0 {
		return n
	}
	for i := int64(0); i < y; i++ {
		n.Mul(n, x)
	}
	return n
}

func numerator(iteration int64, precision uint) *big.Float {
	numerator := big.NewInt(1)
	if iteration%2 != 0 {
		numerator.Mul(numerator, big.NewInt(-1))
	}
	six_k := big.NewInt(6 * iteration)
	six_k = factorial(six_k)
	first_sum_k := big.NewInt(545140134 * iteration)
	first_sum_k.Add(first_sum_k, big.NewInt(13591409))
	numerator.Mul(numerator, six_k)
	numerator.Mul(numerator, first_sum_k)
	float_numerator := new(big.Float).SetInt(numerator).SetPrec(precision)
	return float_numerator
}

func denominator(iteration int64, precision uint) *big.Float {
	three_k_fac := big.NewInt(3 * iteration)
	three_k_fac = factorial(three_k_fac)
	three_k_fac_float := new(big.Float).SetInt(three_k_fac).SetPrec(precision)
	k_fac := big.NewInt(iteration)
	k_fac = factorial(k_fac)
	k_fac_cubed := big.NewInt(1)
	k_fac_cubed.Mul(k_fac, k_fac)
	k_fac_cubed.Mul(k_fac_cubed, k_fac)

	k_fac_cubed_float := new(big.Float).SetInt(k_fac_cubed).SetPrec(precision)
	exponent := 6*iteration + 3

	pow1 := big.NewInt(640320)
	pow1 = power(pow1, exponent)
	denom := new(big.Float).SetInt(pow1).SetPrec(precision)
	denom.Sqrt(denom)

	denom.Mul(denom, three_k_fac_float)
	denom.Mul(denom, k_fac_cubed_float)
	return denom
}

// ComputePi calculates pi to a given number of decimal places using the Chudnovsky algorithm
func computePi(digits uint) *big.Float {

	// Set the precision to ensure accuracy
	bigPrec := digits * 4
	sum := big.NewFloat(0).SetPrec(bigPrec)
	one := big.NewFloat(1).SetPrec(bigPrec)
	newpi := big.NewFloat(1).SetPrec(bigPrec)
	pi := big.NewFloat(0).SetPrec(bigPrec)
	twelve := big.NewFloat(12).SetPrec(bigPrec)
	var i int64 = 0
	for {
		pi.Set(newpi)
		numerator := numerator(i, bigPrec)
		denominator := denominator(i, bigPrec)
		term := big.NewFloat(0).SetPrec(bigPrec)
		sum.Add(sum, term.Quo(numerator, denominator))
		newpi.Quo(one, sum)
		newpi.Quo(newpi, twelve)
		if pi.Cmp(newpi) == 0 {
			break
		}
		i++
		fmt.Printf("Iteration %d\n", i)
	}
	fmt.Printf("It took %d iterations to calculate pi to %d places.\n", i+1, digits)
	return pi
}

func main() {
	// Desired number of decimal places
	digits := uint(1000)

	// Compute pi
	pi := computePi(digits)

	// Print pi to the desired precision
	fmt.Printf("Pi to %d decimal places:\n%s\n", digits, pi.Text('f', int(digits)+5))
}

package main

import (
	"fmt"
	"math/big"
)

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}

func main() {
	precision := uint(1000)
	digits := precision * 4
	one := big.NewFloat(1).SetPrec(digits)
	zero := big.NewFloat(0).SetPrec(digits)
	euler := big.NewFloat(0).SetPrec(digits)
	new_euler := big.NewFloat(0).SetPrec(digits)
	quotient := big.NewFloat(0).SetPrec(digits)
	sum := big.NewFloat(0).SetPrec(digits)
	iterator := 0
	for {
		euler.Set(new_euler).SetPrec(digits)

		denom := factorial(big.NewInt(int64(iterator)))

		quotient.Quo(one, new(big.Float).SetInt(denom).SetPrec(digits))
		sum.Add(sum, quotient)
		new_euler = sum.Add(sum, zero)
		if new_euler.Cmp(euler) == 0 {
			break
		}
		iterator++
		// fmt.Printf("e to %d decimal places:\n%s\n", precision, new_euler.Text('f', int(precision)+5))
	}
	fmt.Printf("It took a value of %d to get %d precision.\n", iterator, precision)
	fmt.Printf("e to %d decimal places:\n%s\n", precision, new_euler.Text('f', int(precision)+5))
}

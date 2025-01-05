package main

import (
	"math/big"
)

func AddBinary(a string, b string) string {
	aBigInt := new(big.Int)
	aBigInt, _ = aBigInt.SetString(a, 2)

	bBigInt := new(big.Int)
	bBigInt, _ = bBigInt.SetString(b, 2)

	sum := new(big.Int)

	sum.Add(aBigInt, bBigInt)

	return sum.Text(2)
}

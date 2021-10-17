package mathematics

import "math/big"

type Arithmetizable interface {
	Add(addings ...Arithmetizable) (Arithmetizable, error)
	Substract(substractings ...Arithmetizable) (Arithmetizable, error)
	Multiply(factors ...Arithmetizable) (Arithmetizable, error)
	Divide(divisors ...Arithmetizable) (Arithmetizable, error)
	Equals(arithmetizable Arithmetizable) bool
	String() string
	GetClientAsBigInt() big.Int
	PartsAsString() (string, string)
}

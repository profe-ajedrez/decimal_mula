package mathematics

import "math/big"

type Arithmetizable interface {
	Add(addings ...Arithmetizable) Arithmetizable
	Substract(substractings ...Arithmetizable) Arithmetizable
	Multiply(factors ...Arithmetizable) Arithmetizable
	Divide(divisors ...Arithmetizable) Arithmetizable
	Equals(arithmetizable Arithmetizable) bool
	Gt(arithmetizable Arithmetizable) bool
	Lt(arithmetizable Arithmetizable) bool
	Gte(arithmetizable Arithmetizable) bool
	Lte(arithmetizable Arithmetizable) bool
	Abs() Arithmetizable
	String() string
	GetClientAsBigInt() big.Int
	PartsAsString() (string, string)
	GetClient() interface{}
	Clone() (Arithmetizable, error)
}

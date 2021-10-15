package mathematics

import "math/big"

type Arithmetizable interface {
	Add(adding Arithmetizable) (Arithmetizable, error)
	Substract(substracting Arithmetizable) (Arithmetizable, error)
	Multiply(factor Arithmetizable) (Arithmetizable, error)
	Divide(divisor Arithmetizable) (Arithmetizable, error)
	Equals(arithmetizable Arithmetizable) bool
	String() string
	GetClientAsBigInt() big.Int
	GetPrecision() int
}

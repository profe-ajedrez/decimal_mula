package mathematics

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"
)

var Precision int = 24
var bigZero = DecimalMulaNeutroAditivo().value
var inverso = DecimalMulaUnidadInversa()

type decimalMula struct {
	value big.Int
}

type operation func(a1 *big.Int, a2 decimalMula) *big.Int

func IsDecimalMula(candidate interface{}) bool {
	return strings.ToUpper(reflect.TypeOf(candidate).String()) == "MATHEMATICS.DECIMALMULA"
}

func DecimalMulaFromString(value string) (decimalMula, error) {
	return dmFromString(value, Precision)
}

func DecimalMulaFromInt(value int) (decimalMula, error) {
	return DecimalMulaFromString(fmt.Sprintf("%d.0", value))
}

func DecimalMulaFromBigInt(value big.Int) (decimalMula, error) {
	return DecimalMulaFromString(value.String())
}

func DecimalMulaNeutroAditivo() decimalMula {
	dm, _ := DecimalMulaFromInt(0)
	return dm
}

func DecimalMulaNeutroMultiplicativo() decimalMula {
	dm, _ := DecimalMulaFromInt(1)
	return dm
}

func DecimalMulaUnidadInversa() decimalMula {
	dm, _ := DecimalMulaFromInt(-1)
	return dm
}

func (d *decimalMula) String() string {
	s := d.value.Text(10)
	integerPart, decimalPart := getPartsFromString(s, Precision)
	return integerPart + "." + decimalPart
}

func (d *decimalMula) Clone() (decimalMula, error) {
	dm, err := DecimalMulaFromString(d.String())
	return dm, err
}

func (d *decimalMula) Add(addings ...decimalMula) *decimalMula {
	a1 := paramLoop(*d, addings, func(a1 *big.Int, a2 decimalMula) *big.Int {
		v := a2.value
		return a1.Add(a1, &v)
	})
	dm := defaultDecimalMula(a1)
	return &dm
}

func (d *decimalMula) Substract(substractings ...decimalMula) *decimalMula {
	a1 := paramLoop(*d, substractings, func(a1 *big.Int, a2 decimalMula) *big.Int {
		v := a2.value
		return a1.Sub(a1, &v)
	})
	dm := defaultDecimalMula(a1)
	return &dm
}

func (d *decimalMula) Multiply(factors ...decimalMula) *decimalMula {
	multiPrecision := Precision * len(factors)
	a1 := paramLoop(*d, factors, func(a1 *big.Int, a2 decimalMula) *big.Int {
		v := a2.value
		a1.Mul(a1, &v)
		return a1
	})

	bigPrecision, err := stringToBigInt(buildFactor(multiPrecision))
	if err == nil {
		dm := defaultDecimalMula(*a1.Div(&a1, bigPrecision))
		return &dm
	}
	panic(err)
}

func (d *decimalMula) Divide(divisors ...decimalMula) *decimalMula {
	factor, err := buildFactorAsBigInt(Precision)
	if err == nil {
		q := paramLoop(*d, divisors, func(a1 *big.Int, a2 decimalMula) *big.Int {
			v := a2.value
			a1.Mul(a1, &factor)
			a1.Div(a1, &v)
			return a1
		})
		iPart, dPart := getPartsFromString(q.Text(10), Precision)
		dm, err := DecimalMulaFromString(iPart + "." + dPart)
		if err == nil {
			return &dm
		}
	}
	panic(err)
}

func (d *decimalMula) Equals(dm decimalMula) bool {
	return d.String() == dm.String()
}

func (d *decimalMula) AdditiveInverse() *decimalMula {
	return d.Multiply(inverso)
}

func (d *decimalMula) MultiplicativeInverse() *decimalMula {
	one := DecimalMulaNeutroMultiplicativo()
	return one.Divide(*d)
}

func (d *decimalMula) PartsAsString() (string, string) {
	iPart, dPart := getPartsFromString(d.String(), Precision)
	return iPart, strings.TrimRight(dPart, "0")
}

func (d *decimalMula) GetClientAsBigInt() big.Int {
	return d.value
}

func (d *decimalMula) Abs() (decimalMula, error) {
	if d.value.Cmp(&bigZero) == -1 {
		absized := d.AdditiveInverse()
		return *absized, nil
	}
	return d.Clone()
}

func (d *decimalMula) Lt(dm decimalMula) bool {
	client := dm.value
	return d.value.Cmp(&client) == -1
}

func (d *decimalMula) Gt(dm decimalMula) bool {
	client := dm.value
	return d.value.Cmp(&client) == 1
}

func (d *decimalMula) Lte(dm decimalMula) bool {
	return d.Lt(dm) || d.Equals(dm)
}

func (d *decimalMula) Gte(dm decimalMula) bool {
	return d.Gt(dm) || d.Equals(dm)
}

func paramLoop(dm decimalMula, params []decimalMula, operatable operation) big.Int {
	a1 := new(big.Int).Set(&dm.value)
	for _, param := range params {
		a1 = operatable(a1, param)
	}
	return *a1
}

func dmFromString(value string, precision int) (decimalMula, error) {
	integerPart, decimalPart := getParts(value)
	iLen := len(integerPart)
	strNumber := integerPart + decimalPart
	strNumber = strNumber + strings.Repeat("0", ImaxOf(precision-len(strNumber)+iLen, 0))
	strNumber = strNumber[0 : precision+iLen]
	v, err := stringToBigInt(strNumber)
	if err == nil {
		return defaultDecimalMula(*v), nil
	}
	return emptyDecimalMula(), err
}

func defaultDecimalMula(v big.Int) decimalMula {
	return decimalMula{
		value: v,
	}
}

func emptyDecimalMula() decimalMula {
	return decimalMula{
		value: big.Int{},
	}
}

func getPartsFromString(s string, precision int) (string, string) {
	var decimalPart, integerPart string

	separation := len(s) - precision
	if separation > 0 {
		decimalPart = s[separation:]
		integerPart = s[0:separation]
	} else {
		decimalPart = strings.Repeat("0", precision-len(s)) + s
		integerPart = "0"
	}
	if integerPart == "" || integerPart == "-" {
		integerPart += "0"
	}
	return integerPart, decimalPart
}

func getParts(value string) (string, string) {
	var integerPart, decimalPart string
	parts := strings.Split(value, ".")

	l := len(parts)
	if l == 1 || (l == 2 && parts[1] == "") {
		decimalPart = "0"
	} else {
		decimalPart = parts[1]
	}

	integerPart = parts[0]
	if integerPart == "" || integerPart == "-" {
		integerPart += "0"
	}

	return integerPart, decimalPart
}

func stringToBigInt(value string) (*big.Int, error) {
	bint, ok := new(big.Int).SetString(value, 10)
	if ok {
		return bint, nil
	}

	return nil, newDecimalMulaError(fmt.Sprintf("Can't convert %s to big.Int", value))
}

func buildFactor(prec int) string {
	return "1" + strings.Repeat("0", prec)
}

func buildFactorAsBigInt(prec int) (big.Int, error) {
	bint, err := stringToBigInt(buildFactor(prec))
	return *bint, err
}

// Temporalmente comentado hasta que se implemente un método de redondeo
//func bigIntFromInt(value int) big.Int {
//	return *big.NewInt(int64(value))
//}

type decimalMulaError struct {
	msg string
}

func newDecimalMulaError(msg string) decimalMulaError {
	return decimalMulaError{
		msg: msg,
	}
}

func (e decimalMulaError) Error() string {
	return e.msg
}

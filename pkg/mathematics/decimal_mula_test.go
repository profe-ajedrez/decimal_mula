package mathematics

import (
	"testing"
)

func TestDecimalMulaCreation(t *testing.T) {
	Precision = 32
	dm1, _ := DecimalMulaFromString("1.123456789012345678901235")
	dm2, _ := dm1.Clone()
	dm1 = *dm2
}

func TestDecimalMulaAdd(t *testing.T) {
	Precision = 32
	dm1, _ := DecimalMulaFromString("1.123456789012345678901235")
	dm2, _ := DecimalMulaFromString("-1.1")
	dm3, _ := DecimalMulaFromString("1.202")
	dm4, _ := DecimalMulaFromString("1.61")

	sum := dm1.Add(dm2, dm3, dm4)
	want := "2.83545678901234567890123500000000"

	if want != sum.String() {
		t.Errorf("Expected %v, got %v at decimalMula.Add()", want, sum.String())
	}
}

func TestDecimalMulaSubstract(t *testing.T) {
	Precision = 32
	dm1, _ := DecimalMulaFromString("11.123456789012345678901235")
	dm2, _ := DecimalMulaFromString("1.1")
	dm3, _ := DecimalMulaFromString("-1.202")
	dm4, _ := DecimalMulaFromString("1.61")
	dm5, _ := DecimalMulaFromString("0.8989")

	sub := dm1.Substract(dm2, dm3, dm4, dm5)
	want := "8.71655678901234567890123500000000"

	if want != sub.String() {
		t.Errorf("Expected %v, got %v at decimalMula.Substract()", want, sub.String())
	}
}

func TestDecimalMulaMultiply(t *testing.T) {
	Precision = 33
	dm1, _ := DecimalMulaFromString("11.123456789012345678901235")
	dm2, _ := DecimalMulaFromString("1.1")
	dm3, _ := DecimalMulaFromString("-1.202")
	dm4, _ := DecimalMulaFromString("1.61")
	dm5, _ := DecimalMulaFromString("0.8989")

	mul := dm1.Multiply(dm2, dm3, dm4, dm5)
	want := "-21.285025820142995598034300386656993"

	if want != mul.String() {
		t.Errorf("Expected %v, got %v at decimalMula.Multiply()", want, mul.String())
	}
}

func TestDecimalMulaDivide(t *testing.T) {
	Precision = 32
	dm1, _ := DecimalMulaFromString("11.123456789012345678901235")
	dm2, _ := DecimalMulaFromString("1.1")
	dm3, _ := DecimalMulaFromString("-1.202")
	dm4, _ := DecimalMulaFromString("1.61")
	dm5, _ := DecimalMulaFromString("0.8989")

	divi := dm1.Divide(dm2, dm3, dm4, dm5)
	want := "-5.81306745796531999297392368402250"

	if want != divi.String() {
		t.Errorf("Expected %v, got %v at decimalMula.Multiply()", want, divi.String())
	}
}

func TestDecimalMulaEquals(t *testing.T) {
	Precision = 32
	dm1, _ := DecimalMulaFromString("11.123456789012345678901235")
	dm2, _ := DecimalMulaFromString("11.123456789012345678901235")
	dm3, _ := DecimalMulaFromString("1.1")

	if !dm1.Equals(dm2) {
		t.Errorf("Expected %v to be equals with %v at decimalMula.Equals()", dm1.String(), dm2.String())
	}

	if dm1.Equals(dm3) {
		t.Errorf("Expected %v to be distinct with %v at decimalMula.Equals()", dm1.String(), dm3.String())
	}
}

func BenchmarkCreateDecimalMula(b *testing.B) {
	b.ResetTimer()
	var d decimalMula
	for n := 0; n < 100000; n++ {
		d, _ = DecimalMulaFromString("1.123456789012345678901235")
	}
	d.Add(d, d)
}

func BenchmarkAddDecimalMula(b *testing.B) {
	b.ResetTimer()
	var d3 decimalMula
	d1, _ := DecimalMulaFromString("1.123456789012345678901235")
	d2, _ := DecimalMulaFromString("3457891.4123456789012345678901235")
	for n := 0; n < 100000; n++ {
		d3 = *d2.Add(d1)
	}
	d3.Add(d3, d3)
}

package mathematics

import (
	"testing"
)

func TestDecimalMulaArithmetizableCreation(t *testing.T) {
	Precision = 32
	dma1, _ := DMArithmetizableFromString("1.123456789012345678901235")
	dma2, _ := dma1.Clone()
	dma1 = dma2
}

func TestDecimalArithmetizableMulaAdd(t *testing.T) {
	dm1, _ := DMArithmetizableFromString("1.123456789012345678901235")
	dm2, _ := DMArithmetizableFromString("-1.1")
	dm3, _ := DMArithmetizableFromString("1.202")
	dm4, _ := DMArithmetizableFromString("1.61")

	sum := dm1.Add(dm2, dm3, dm4)
	want := "2.83545678901234567890123500000000"

	if want != sum.String() {
		t.Errorf("Expected %v, got %v at decimalMulaArithmetizable.Add()", want, sum.String())
	}
}

func TestDecimalMulaArithmetizableSubstract(t *testing.T) {
	dm1, _ := DMArithmetizableFromString("11.123456789012345678901235")
	dm2, _ := DMArithmetizableFromString("1.1")
	dm3, _ := DMArithmetizableFromString("-1.202")
	dm4, _ := DMArithmetizableFromString("1.61")
	dm5, _ := DMArithmetizableFromString("0.8989")

	sub := dm1.Substract(dm2, dm3, dm4, dm5)
	want := "8.71655678901234567890123500000000"

	if want != sub.String() {
		t.Errorf("Expected %v, got %v at decimalMulaArithmetizable.Substract()", want, sub.String())
	}
}

func TestDecimalMulaArithmetizableMultiply(t *testing.T) {
	Precision = 33
	dm1, _ := DMArithmetizableFromString("11.123456789012345678901235")
	dm2, _ := DMArithmetizableFromString("1.1")
	dm3, _ := DMArithmetizableFromString("-1.202")
	dm4, _ := DMArithmetizableFromString("1.61")
	dm5, _ := DMArithmetizableFromString("0.8989")

	mul := dm1.Multiply(dm2, dm3, dm4, dm5)
	want := "-21.285025820142995598034300386656993"

	if want != mul.String() {
		t.Errorf("Expected %v, got %v at decimalMulaArithmetizable.Multiply()", want, mul.String())
	}
}

func TestDecimalMulaArithmetizableDivide(t *testing.T) {
	Precision = 32
	dm1, _ := DMArithmetizableFromString("11.123456789012345678901235")
	dm2, _ := DMArithmetizableFromString("1.1")
	dm3, _ := DMArithmetizableFromString("-1.202")
	dm4, _ := DMArithmetizableFromString("1.61")
	dm5, _ := DMArithmetizableFromString("0.8989")

	divi := dm1.Divide(dm2, dm3, dm4, dm5)
	want := "-5.81306745796531999297392368402250"

	if want != divi.String() {
		t.Errorf("Expected %v, got %v at decimalMulaArithmetizable.Multiply()", want, divi.String())
	}
}

func TestDecimalMulaArithmetizableCompare(t *testing.T) {
	Precision = 32
	dm1, _ := DMArithmetizableFromString("11.123456789012345678901235")
	dm2, _ := DMArithmetizableFromString("11.123456789012345678901235")
	dm3, _ := DMArithmetizableFromString("1.1")

	if !dm1.Equals(dm2) {
		t.Errorf("Expected %v to be equals with %v at decimalMulaArithmetizable.Equals()", dm1.String(), dm2.String())
	}

	if dm1.Equals(dm3) {
		t.Errorf("Expected %v to be distinct with %v at decidecimalMulaArithmetizablemalMula.Equals()", dm1.String(), dm3.String())
	}

	if !dm1.Gt(dm3) {
		t.Errorf("Expected %v to be greater than %v at decimalMulaArithmetizable.Gt()", dm1.String(), dm3.String())
	}

	if !dm3.Lt(dm1) {
		t.Errorf("Expected %v to be lesser than %v at decimalMulaArithmetizable.Lt()", dm3.String(), dm1.String())
	}

	if !dm1.Gte(dm2) {
		t.Errorf("Expected %v to be greater or equal than %v at decimalMulaArithmetizable.Gte()", dm1.String(), dm2.String())
	}

	if !dm1.Lte(dm2) {
		t.Errorf("Expected %v to be lesser or equal than %v at decimalMulaArithmetizable.Lte()", dm1.String(), dm2.String())
	}

}

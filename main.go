package main

import (
	"fmt"
	"math/big"

	"github.com/profe-ajedrez/decimal_mula/pkg/mathematics"
)

func main() {
	d1, _ := mathematics.DecimalMulaFromString("-0.123")
	d2, _ := mathematics.DecimalMulaFromString("-1.234")
	d3, _ := mathematics.DecimalMulaFromString("2.345")
	d4, _ := mathematics.DecimalMulaFromString("3.456789")

	b1, _ := mathematics.DecimalMulaFromInt(4)
	b2, _ := mathematics.DecimalMulaFromBigInt(*big.NewInt(int64(5)))

	fmt.Println(d1.Add(d2, d3, d4, b1, b2).String())
	fmt.Println(d4.Substract(d3, d2))

	aa := d3.Divide(d4)
	fmt.Println(aa.String())

	//b3 := d2.Add(d4)
	s1 := d1.String()
	s2 := d2.String()
	s3 := d3.String()
	s4 := b1.String()
	s5 := b2.String()
	s6 := d4.String()
	//s7 := b3.String()

	fmt.Println("s1", s1)
	fmt.Println("s2", s2)
	fmt.Println("s3", s3)
	fmt.Println("s4", s4)
	fmt.Println("s5", s5)
	fmt.Println("s6", s6)

	dm1, _ := mathematics.DecimalMulaFromString("1.123456789012345678901235")
	dm2, _ := mathematics.DecimalMulaFromString("1.1")
	dm3, _ := mathematics.DecimalMulaFromString("1.2")
	dm4, _ := mathematics.DecimalMulaFromString("1.6")
	dm5, _ := mathematics.DecimalMulaFromString("10.1234567890123456789012350")
	dm6, _ := mathematics.DecimalMulaFromString("1.10")
	dm7, _ := mathematics.DecimalMulaFromString("1.20")
	dm8, _ := mathematics.DecimalMulaFromString("1.60")
	dm9, _ := mathematics.DecimalMulaFromString("7.98464675343434")

	as1 := dm1.Add(dm2, dm3, dm4)
	ar1, _ := dm5.Clone()
	ar11 := ar1.Substract(dm6, dm7, dm8)
	ar2, _ := dm5.Clone()
	ar22 := ar2.Substract(dm6, dm7, dm8, dm9)

	//want1 := "5.023456789012345678901235"
	//want2 := "6.223456789012345678901235"
	//want3 := "-1.761189964421994321098765"

	//gotSum := s1.String()
	//gotSub := r11.String()
	//gotSub2 := r22.String()

	fmt.Println(as1.String())
	fmt.Println(ar1.String())
	fmt.Println(ar11.String())
	fmt.Println(ar2.String())
	fmt.Println(ar22.String())

}

package main

import (
	"fmt"
)

type ErrDivideByZero float64

func (e ErrDivideByZero) Error() string {
	return fmt.Sprintf("cannot divide %v by zero:", float64(e))
	//return fmt.Sprintf("cannot Sqrt negative number: %v", e)//this blows out the stack with recursive calls to Error() because e is type
	//ErrNegativeSqrt and this function gets called over and over
}

func Quotient(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, ErrDivideByZero(dividend)
	}
	return dividend / divisor, nil
}

func main() {
	fmt.Println(Quotient(2, 2000000))
	fmt.Println(Quotient(2, 0))
}

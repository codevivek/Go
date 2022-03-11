/*44 –	You just won a million dollars! You decide to invest $500,000 in a term
deposit for a period of 5 years. The annual interest rate is 10%, and the
interest is added to the principal sum each year (compound interest).
How much will your savings be worth in 5 years?

a)	Make the algorithm with specified number (internal data).
b)	Generalize for any amount, any duration, and any interest rate. */


package main

import (
	"fmt"
	"strconv"
)

func main(){
	var annualInterestRate float32 = .1
	var captialInvestement float32 =500000
	termInYears :=5

	i :=1
	for i<=termInYears {
		captialInvestement = captialInvestement + (captialInvestement*annualInterestRate)
		s := fmt.Sprintf("%.0f", captialInvestement) //
		fmt.Println( " year : " + strconv.Itoa(i) + " = " + s + "$" )
		i += 1
	}


}




package mortgage

import "fmt"

func CalculateMortgage(payload *PayloadMortgage) int {
	fmt.Printf("payload %#v", payload.AnnualInterestRate)
	return 0
}

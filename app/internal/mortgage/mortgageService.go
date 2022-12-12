package mortgage

import (
	"github.com/shopspring/decimal"
)

const (
	BiWeeklyNumOfPayments      = 24
	BiWeeklyAccelNumOfPayments = 26
	MonthlyNumOfPayments       = 12
)

// CalculateMortgage Function will calculate mortgage payments
// calculation off compared to wolfram alpha vs www.ratehub.ca
// https://www.wolframalpha.com/input?i2d=true&i=Divide%5B278370*Divide%5B0.05%2C12%5D*Power%5B%5C%2840%291+%2B+Divide%5B0.05%2C12%5D%5C%2841%29%2C300%5D%2CPower%5B%5C%2840%291%2BDivide%5B0.05%2C12%5D%5C%2841%29%2C300%5D-1%5D
func CalculateMortgage(payload *PayloadMortgage) float64 {

	var numberOfPayments int
	interestRate := decimal.NewFromFloat(payload.AnnualInterestRate / 100)

	var yearlyNumberOfPayments int
	switch payload.PaymentSchedule {
	case Monthly:
		numberOfPayments = MonthlyNumOfPayments * payload.Period
		yearlyNumberOfPayments = MonthlyNumOfPayments
	case BiWeekly:
		numberOfPayments = BiWeeklyNumOfPayments * payload.Period
		yearlyNumberOfPayments = BiWeeklyNumOfPayments
	case BiWeeklyAccel:
		numberOfPayments = BiWeeklyAccelNumOfPayments * payload.Period
		yearlyNumberOfPayments = BiWeeklyAccelNumOfPayments
	}

	cmhcRate := CalculateCMHC(decimal.NewFromFloat(payload.DownPayment), decimal.NewFromFloat(payload.PropertyPrice))
	cmhcTotal := decimal.NewFromFloat(payload.PropertyPrice - payload.DownPayment).Mul(cmhcRate)
	monthlyInterestRate := interestRate.Div(decimal.NewFromInt(int64(yearlyNumberOfPayments)))
	principle := decimal.NewFromFloat(payload.PropertyPrice - payload.DownPayment).Add(cmhcTotal)

	interestRatePlusOne := monthlyInterestRate.Add(decimal.NewFromInt(1))
	interestRatePlusOneToPower := interestRatePlusOne.Pow(decimal.NewFromInt(int64(numberOfPayments)))

	equationTop := principle.Mul(monthlyInterestRate.Mul(interestRatePlusOneToPower))
	equationBottom := interestRatePlusOneToPower.Sub(decimal.NewFromInt(1))

	mortgagePayments, _ := equationTop.Div(equationBottom).Float64()

	return mortgagePayments
}

func CalculateCMHC(downPayment decimal.Decimal, price decimal.Decimal) decimal.Decimal {
	downPaymentPercentage := downPayment.Div(price)
	if downPaymentPercentage.GreaterThanOrEqual(decimal.NewFromFloat(0.00)) &&
		downPaymentPercentage.LessThanOrEqual(decimal.NewFromFloat(9.99/100)) {
		return decimal.NewFromFloat(0.04)
	} else if downPaymentPercentage.GreaterThanOrEqual(decimal.NewFromFloat(10.00/100)) &&
		downPaymentPercentage.LessThanOrEqual(decimal.NewFromFloat(14.99/100)) {
		return decimal.NewFromFloat(0.031)
	} else if downPaymentPercentage.GreaterThanOrEqual(decimal.NewFromFloat(15.00/100)) &&
		downPaymentPercentage.LessThanOrEqual(decimal.NewFromFloat(19.00/100)) {
		return decimal.NewFromFloat(0.028)
	} else {
		return decimal.NewFromFloat(0.00)
	}
}

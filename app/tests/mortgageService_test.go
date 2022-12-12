package tests

import (
	"fmt"
	"github.com/shopspring/decimal"
	"quoter/internal/mortgage"
	"testing"
)

func TestCalculateMortgageBiWeek(t *testing.T) {
	expected := 813.2369417096032044
	payload := mortgage.PayloadMortgage{
		PropertyPrice:      300000,
		DownPayment:        30000,
		AnnualInterestRate: 5.00,
		Period:             25,
		PaymentSchedule:    "bi-weekly",
	}
	actual := mortgage.CalculateMortgage(&payload)

	if expected != actual {
		t.Errorf("expected '%f' but got '%f'", expected, actual)
	}
}

func TestCalculateMortgageMonthly(t *testing.T) {
	expected := 1627.3232985457677374
	payload := mortgage.PayloadMortgage{
		PropertyPrice:      300000,
		DownPayment:        30000,
		AnnualInterestRate: 5.00,
		Period:             25,
		PaymentSchedule:    "monthly",
	}
	actual := mortgage.CalculateMortgage(&payload)
	fmt.Println(actual)

	if expected != actual {
		t.Errorf("expected '%f' but got '%f'", expected, actual)
	}
}

func TestCalculateMortgageBiWeekAccel(t *testing.T) {
	expected := 750.6500901675546442
	payload := mortgage.PayloadMortgage{
		PropertyPrice:      300000,
		DownPayment:        30000,
		AnnualInterestRate: 5.00,
		Period:             25,
		PaymentSchedule:    "bi-weekly-accel",
	}
	actual := mortgage.CalculateMortgage(&payload)

	if expected != actual {
		t.Errorf("expected '%f' but got '%f'", expected, actual)
	}
}

func TestCalculateCMHCTier1(t *testing.T) {
	expected := decimal.NewFromFloat(0.04)
	downPayment := decimal.NewFromFloat(25000)
	price := decimal.NewFromFloat(300000)

	actualCMHC := mortgage.CalculateCMHC(downPayment, price)

	if !expected.Equal(actualCMHC) {
		t.Errorf("expected '%d' but got '%d'", expected, actualCMHC)
	}
}

func TestCalculateCMHCTier2(t *testing.T) {
	expected := decimal.NewFromFloat(0.031)
	downPayment := decimal.NewFromFloat(30000)
	price := decimal.NewFromFloat(300000)

	actualCMHC := mortgage.CalculateCMHC(downPayment, price)
	fmt.Println(actualCMHC)
	if !expected.Equal(actualCMHC) {
		t.Errorf("expected '%d' but got '%d'", expected, actualCMHC)
	}
}

func TestCalculateCMHCTier3(t *testing.T) {
	expected := decimal.NewFromFloat(0.028)
	downPayment := decimal.NewFromFloat(45000)
	price := decimal.NewFromFloat(300000)

	actualCMHC := mortgage.CalculateCMHC(downPayment, price)
	fmt.Println(actualCMHC)
	if !expected.Equal(actualCMHC) {
		t.Errorf("expected '%d' but got '%d'", expected, actualCMHC)
	}
}

func TestCalculateCMHCTier4(t *testing.T) {
	expected := decimal.NewFromFloat(0.00)
	downPayment := decimal.NewFromFloat(60000)
	price := decimal.NewFromFloat(300000)

	actualCMHC := mortgage.CalculateCMHC(downPayment, price)
	fmt.Println(actualCMHC)
	if !expected.Equal(actualCMHC) {
		t.Errorf("expected '%d' but got '%d'", expected, actualCMHC)
	}
}

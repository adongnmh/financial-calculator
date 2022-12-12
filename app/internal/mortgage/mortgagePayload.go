package mortgage

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
)

const (
	BiWeekly      = "bi-weekly"
	BiWeeklyAccel = "bi-weekly-accel"
	Monthly       = "monthly"
)

type PayloadMortgage struct {
	PropertyPrice      float64 `json:"property_price" validate:"gte=1"`
	DownPayment        float64 `json:"down_payment" validate:"gte=0"`
	AnnualInterestRate float64 `json:"annual_interest_rate" validate:"gte=0"`
	Period             int     `json:"period" validate:"required"`
	PaymentSchedule    string  `json:"payment_schedule" validate:"required,schedule"`
}

func (mr *PayloadMortgage) decodePayload(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(mr)
	if err != nil {
		return err
	}
	return nil
}

func (mr *PayloadMortgage) validateFields(payload *PayloadMortgage) error {
	validate := validator.New()
	err := validate.RegisterValidation("schedule", validatePaymentSchedule)
	if err != nil {
		return err
	}
	return validate.Struct(payload)
}

func validatePaymentSchedule(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case BiWeekly, BiWeeklyAccel, Monthly:
		return true
	}
	return false
}

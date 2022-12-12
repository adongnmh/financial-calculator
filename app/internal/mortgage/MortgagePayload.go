package mortgage

import (
	"encoding/json"
	"io"
)

const (
	BiWeekly      string = "bi-weekly"
	BiWeeklyAccel        = "bi-weekly-accelerated"
	Monthly              = "monthly"
)

type Payload struct {
	PropertyPrice      float64 `json:"property_price"`
	DownPayment        float64 `json:"down_payment"`
	AnnualInterestRate float64 `json:"annual_interest_rate"`
	Period             int64   `json:"period"`
	PaymentSchedule    string  `json:"payment_schedule"`
}

func (mr *Payload) decodePayload(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(mr)
	if err != nil {
		return err
	}
	return nil
}

### Canada Mortage Calculator

- Simple REST API that calculates the mortgage payments 

#### To run server:
docker-compose up --build
or 
go build main.go

serve on localhost:3000

localhost:3000/api/calculate-mortgage-payments
Sample Payload:
```json
{
	"property_price": 300000,
	"down_payment": 30000,
	"annual_interest_rate": 5.00,
	"Period": 25,
	"payment_schedule": "monthly/bi-weekly/bi-weekly-accel"
}
```

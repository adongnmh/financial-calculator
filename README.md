### Canada Mortage Calculator

- Simple REST API that calculates the mortgage payments 

#### Notes:
- Going to use Gorilla mux since I am most familiar with the package for handling routes. I know that the package is being archived - but the package is in a stable state
- Replacements for Gorilla/mux would be chi or maybe use a full web framework, but I prefer not to overcomplicate the application from the start.
- I am using a docker-compose.yaml from a starter project to get the API up and running locally faster.
- The values that I got from https://www.ratehub.ca/british-columbia-mortgage-calculator differs from actual calculations. Please refer to https://www.wolframalpha.com/input?i2d=true&i=Divide%5B278370*Divide%5B0.05%2C24%5D*Power%5B%5C%2840%291+%2B+Divide%5B0.05%2C24%5D%5C%2841%29%2C650%5D%2CPower%5B%5C%2840%291%2BDivide%5B0.05%2C24%5D%5C%2841%29%2C650%5D-1%5D 
- Please feel free to submit comments on the PR for questions, concerns, and suggestions.
#### TODO + Future Considerations
- Add a postgres DB layer - for eventual DB saving.
- For the sake of finishing the assignment in 90min time frame. I did not include any integrated tests and I rely on unit tests to ensure that functionality was working. If there was more time, or I had time after work I would implement this. 
- Would like to streamline the error messages better in the validation of the mortgage payload.
- More custom validation for the Mortgage payload - ie handle ranges, make sure values make sense, etc

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

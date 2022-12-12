### Quoter Coding Challenge

- Simple REST API that calculates the mortgage payments 

#### Notes:
- Going to use Gorilla mux since I am most familiar with the package for handling routes. I know that the package is being archived - but the package is in a stable state
- Replacements for Gorilla/mux would be chi or maybe use a full web framework, but I prefer not to overcomplicate the application from the start.
- I am using a docker-compose.yaml from a starter project to get the API up and running locally faster.
- Please feel free to submit comments on the PR for questions, concerns, and suggestions.
#### TODO + Future Considerations
- Add a postgres DB layer - for eventual DB saving.
- For the sake of finishing the assignment in 90min time frame. I did not include any integrated tests and I rely on unit tests to ensure that functionality was working. If there was more time, or I had time after work I would implement this. 
- Would like to streamline the error messages better in the validation of the mortgage payload.
- More custom validation for the Mortgage payload - ie handle ranges, make sure values make sense, etc
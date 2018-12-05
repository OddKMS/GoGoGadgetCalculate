package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type CalcEvent struct {
	Number1  int    `json:"Number 1"`
	Number2  int    `json:"Number 2"`
	Operator string `json:"Operator"`
}

type CalcResponse struct {
	Result string `json:"Result:"`
}

func HandleLambdaEvent(event CalcEvent) (CalcResponse, error) {
	return CalcResponse{
			//result from calculation
			Result: "test"},
		nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}

//To get this up and running ASAP I'm intentionally foregoing tests and edge cases.
//Welcome to programming horror stories chpt. 1 - "Trusting the user"
func calculate(num1, num2 int, opr string) (result int) {
	switch operator := opr; operator {
	case "+":
	case "add":
		{
			return num1 + num2
		}
	case "-":
	case "sub":
		{
			return num1 - num2
		}
	case "*":
	case "mul":
		{
			return num1 * num2

		}
	case "/":
	case "div":
		{
			//Okay, so I don't trust the user THAT much.
			if num2 == 0 {
				return 0
			}
			return num1 / num2
		}
	default:
		return 0
	}
	return 0
}

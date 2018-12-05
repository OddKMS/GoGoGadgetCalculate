package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type CalcEvent struct {
	Number1  int    `json:"Number1"`
	Number2  int    `json:"Number2"`
	Operator string `json:"Operator"`
}

type CalcResponse struct {
	Result int `json:"Result:"`
}

func HandleLambdaEvent(event CalcEvent) (CalcResponse, error) {
	var calcResult = calculate(event.Number1, event.Number2, event.Operator)

	return CalcResponse{
			Result: calcResult},
		nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}

//Simple calculation of numbers, supporting addition, subtraction, multiplication, and division.
func calculate(num1, num2 int, opr string) (result int) {
	switch operator := opr; operator {
	//Uses fallthrough to synonymous cases
	case "+":
		fallthrough
	case "add":
		{
			return num1 + num2
		}
	case "-":
		fallthrough
	case "sub":
		{
			return num1 - num2
		}
	case "*":
		fallthrough
	case "mul":
		{
			return num1 * num2

		}
	case "/":
		fallthrough
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
}

package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"strings"
)

type CalcEvent struct {
	Number1  int    `json:"Number1"`
	Number2  int    `json:"Number2"`
	Operator string `json:"Operator"`
}

//Includes the input so we can verify the result ourselves if needed.
type CalcResponse struct {
	Number1  int    `json:"Number1"`
	Number2  int    `json:"Number2"`
	Operator string `json:"Operator"`
	Result   int    `json:"Result:"`
}

func HandleLambdaEvent(event CalcEvent) (CalcResponse, error) {
	return CalcResponse{
			Number1:  event.Number1,
			Number2:  event.Number2,
			Operator: event.Operator,
			Result:   calculate(event.Number1, event.Number2, event.Operator)},
		nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}

//Simple calculation of numbers, supporting addition, subtraction, multiplication, and division.
func calculate(num1, num2 int, opr string) (result int) {
	lowerOpr := strings.ToLower(opr)

	switch operator := lowerOpr; operator {
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

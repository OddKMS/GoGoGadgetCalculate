package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"math"
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
		Result: "test" 
	},
	nil
}

func main(){
	lambda.Start(HandleLambdaEvent)
}

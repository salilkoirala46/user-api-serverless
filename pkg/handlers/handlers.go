package handlers

import(
	"awsserverless/pkg/user"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/aws"
)

var ErroorMethodNotAllowed = "Method Not allowed"

func GetUser(){

}

func CreateUser(){

}

func UpdateUser(){

}

func DeleteUser(){

}

func UnhandledMethod()(*events.APIGatewayProxyResponse, error){
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
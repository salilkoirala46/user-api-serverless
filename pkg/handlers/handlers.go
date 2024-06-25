package handlers

import(
	"awsserverless/pkg/user"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/aws"
)

var ErroorMethodNotAllowed = "Method Not allowed"

type ErrorBody Struct {
	ErrorMessage string `json:"error, omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*events.APIGatewayProxyResponse, error){
	email := req.PathParameters["email"]

	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynaClient)

		if err != nil {
			return apiResponse(http.StatusBadRequestt, ErrorBody{aws.String(err.Error())}), nil
		}

		return apiResponse(http.StatusOK, result), nil
	}

	result, err := user.FetchUser(tableName, dynaClient)

	if err != nil {
		return apiResponse(http.StatusBadRequestt, ErrorBody{aws.String(err.Error())}), nil
	}

	return apiResponse(http.StatusOK, result), nil
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*events.APIGatewayProxyResponse, error){

}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*events.APIGatewayProxyResponse, error){

}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*events.APIGatewayProxyResponse, error){

}

func UnhandledMethod()(*events.APIGatewayProxyResponse, error){
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
package user 

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aw-sdk-go/aws"
	"github.com/aws/aws-sdk-go/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	error ErrorFailedToFetchRecord = "Failed to fetch record"
	error ErrorFailedToUnMarshalRecord = "Failed to unmarshal record"
)

type User struct{
	Email string `json:"email"`	
	FirstName string `json:"firstName"`
	lastName string `json:"lastName"`
}

func FetchUser(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*User, error){
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableName),
	}

	result, err := dynaClient.GetItem(input)

	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	item := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)

	if err != nil {
		return nil, errors.New(ErrorFailedToUnMarshalRecord)
	}

	return item, nil
}

func FetchUsers(tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*[]User, error){

	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := dynaClient.Scan(input)

	if err !=nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	item := new([]User)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, item)
}

func CreateUser()(){

}

func UpdateUser()(){

}

func DeleteUser() error{

}
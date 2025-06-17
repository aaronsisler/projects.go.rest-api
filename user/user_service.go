package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserService struct {
	// potentially DB connection, etc.
	ddb *dynamodb.Client
}

func NewUserService(ddb *dynamodb.Client) *UserService {
	return &UserService{
		ddb: ddb,
	}
}

func (s *UserService) GetUsers() ([]User, error) {
	out, err := s.ddb.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String("SERVICES_EVENTS_ADMIN_SERVICE"),
		KeyConditionExpression: aws.String("partitionKey = :pk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "USER"},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	users := make([]User, 0, len(out.Items))
	for _, item := range out.Items {
		user, err := mapItemToUser(item)
		if err != nil {
			// Log and skip the bad item instead of failing the whole request
			fmt.Printf("error mapping item to user: %v\n", err)
			continue
		}
		users = append(users, *user)
	}

	return users, nil
}

func (s *UserService) GetUserByID(id string) (*User, error) {
	out, err := s.ddb.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("SERVICES_EVENTS_ADMIN_SERVICE"),
		Key: map[string]types.AttributeValue{
			"partitionKey": &types.AttributeValueMemberS{Value: "USER"},
			"sortKey":      &types.AttributeValueMemberS{Value: "USER#" + id},
		},
	})

	if err != nil {
		return nil, err
	}

	if out.Item == nil {
		return nil, nil
	}

	user, err := mapItemToUser(out.Item)

	return user, err
}

func mapItemToUser(item map[string]types.AttributeValue) (*User, error) {
	var user User

	err := attributevalue.UnmarshalMap(item, &user)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal DynamoDB item to User: %w", err)
	}

	if v, ok := item["sortKey"].(*types.AttributeValueMemberS); ok {
		user.ID = strings.TrimPrefix(v.Value, "USER#")
	} else {
		return nil, fmt.Errorf("name attribute missing or wrong type")
	}

	return &user, nil
}

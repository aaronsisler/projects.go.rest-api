package user

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
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
	user := &User{}

	// Map "name" (string)
	if v, ok := item["name"].(*types.AttributeValueMemberS); ok {
		user.Name = v.Value
	} else {
		return nil, fmt.Errorf("name attribute missing or wrong type")
	}

	// Map "establishmentIds" (string set)
	if v, ok := item["establishmentIds"].(*types.AttributeValueMemberSS); ok {
		user.EstablishmentIds = v.Value
	} else {
		// Could be missing, so just assign empty slice if absent
		user.EstablishmentIds = []string{}
	}

	return user, nil
}

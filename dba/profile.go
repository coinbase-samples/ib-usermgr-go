package dba

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/coinbase-samples/ib-usermgr-go/model"
)

func (m *Repository) ReadProfile(id string) (model.ProfileResponse, error) {
	var profile model.ProfileResponse

	out, err := m.Svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(m.App.ProfileTableName),
		Key: map[string]types.AttributeValue{
			"UserId": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		return profile, err
	}

	//fmt.Println(&out.Item)
	err = attributevalue.UnmarshalMap(out.Item, &profile)
	if err != nil {
		return profile, err
	}
	//fmt.Println(&profile)

	return profile, nil

}

func (m *Repository) UpdateProfile(id string, updateBody model.UpdateProfileRequest) (model.ProfileResponse, error) {
	var profile model.ProfileResponse

	updateItem, err := attributevalue.MarshalMap(updateBody)

	if err != nil {
		return profile, err
	}

	_, err = m.Svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(m.App.ProfileTableName),
		Item:      updateItem,
	})

	if err != nil {
		return profile, err
	}

	profile = model.ProfileResponse(updateBody)

	return profile, nil
}

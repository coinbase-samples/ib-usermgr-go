/**
 * Copyright 2022 Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dba

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/coinbase-samples/ib-usermgr-go/model"
)

func (m *DynamoRepository) ReadProfile(id string) (model.ProfileResponse, error) {
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

	err = attributevalue.UnmarshalMap(out.Item, &profile)
	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (m *DynamoRepository) UpdateProfile(id string, updateBody model.UpdateProfileRequest) (model.ProfileResponse, error) {
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

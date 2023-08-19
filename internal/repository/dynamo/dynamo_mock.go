package dynamo

import (
	"context"

	ddb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/martinarias-uala/go-validacion/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockDynamoDBClient struct {
	mock.Mock
}

func (m *MockDynamoDBClient) PutItem(ctx context.Context, inp *ddb.PutItemInput, opt ...func(*ddb.Options)) (*ddb.PutItemOutput, error) {
	args := m.Called(ctx, inp, opt)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*ddb.PutItemOutput), nil
}

func (m *MockDynamoDBClient) ExecuteStatement(ctx context.Context, inp *ddb.ExecuteStatementInput, opt ...func(*ddb.Options)) (*ddb.ExecuteStatementOutput, error) {
	args := m.Called(ctx, inp, opt)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*ddb.ExecuteStatementOutput), nil
}

type MockDynamoDBRepository struct {
	mock.Mock
}

func (m *MockDynamoDBRepository) GetShape(shapeType string) ([]models.ShapeData, error) {
	args := m.Called(shapeType)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	return []models.ShapeData{}, nil
}

func (m *MockDynamoDBRepository) CreateItem(shape models.ShapeData) error {
	args := m.Called(shape)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return nil
}

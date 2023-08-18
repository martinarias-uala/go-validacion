package s3

/*
import (
	"context"

	"github.com/Uilobank/uilo-loan-portfolio-purchase/commons/models"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) GetObject(ctx context.Context, inp *s3.GetObjectInput, opt ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	args := m.Called(ctx, inp, opt)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*s3.GetObjectOutput), nil
}

func (m *MockClient) PutObject(ctx context.Context, inp *s3.PutObjectInput, opt ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	args := m.Called(ctx, inp, opt)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*s3.PutObjectOutput), nil
}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetObjectContent(bucket, key string) (string, models.ACTraceItem) {
	args := m.Called(bucket, key)
	if args.Get(1) != nil {
		return "", args.Get(1).(models.ACTraceItem)
	}
	return args.Get(0).(string), models.ACTraceItem{}
}

func (m *MockRepository) PutObject(bucket, bucket_key string, loan []models.LoanPortfolio) models.ACTraceItem {
	args := m.Called(bucket, bucket_key, loan)
	if args.Get(0) != nil {
		return args.Get(0).(models.ACTraceItem)
	}
	return models.ACTraceItem{}
}
*/

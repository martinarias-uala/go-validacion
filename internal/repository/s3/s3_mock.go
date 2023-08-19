package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/martinarias-uala/go-validacion/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockS3Client struct {
	mock.Mock
}

func (m *MockS3Client) PutObject(ctx context.Context, inp *s3.PutObjectInput, opt ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	args := m.Called(ctx, inp, opt)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*s3.PutObjectOutput), nil
}

type MockS3Repository struct {
	mock.Mock
}

func (m *MockS3Repository) PutObject(shape models.ShapeData) error {
	args := m.Called(shape)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return nil
}

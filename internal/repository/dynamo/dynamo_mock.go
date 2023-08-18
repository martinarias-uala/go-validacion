package dynamo

/* import (
	"context"

	ddb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) PutItem(ctx context.Context, inp *ddb.PutItemInput, opt ...func(*ddb.Options)) (*ddb.PutItemOutput, error) {
	args := m.Called(ctx, inp, opt)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*ddb.PutItemOutput), nil
}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateItem(i models.ACTraceItem) error {
	args := m.Called(i)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return nil
}
*/

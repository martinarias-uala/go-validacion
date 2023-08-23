package service

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockHttp struct {
	mock.Mock
}

func (m *MockHttp) Get(url string) (*http.Response, error) {
	args := m.Called(url)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*http.Response), nil
}

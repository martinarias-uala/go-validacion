package shapes

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockShapesController struct {
	mock.Mock
}

func (m *MockShapesController) GetShapes(ctx *gin.Context) {
	m.Called(ctx)
}
func (m *MockShapesController) CreateShape(ctx *gin.Context) {
	m.Called(ctx)
}

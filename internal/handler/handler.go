package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/martinarias-uala/go-validacion/internal/shapes"
)

type Handler struct {
	sc shapes.IShapesController
}

var ginLambda *ginadapter.GinLambda

func (h Handler) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		r := gin.Default()
		r.GET("/shapes", h.sc.GetShape)
		r.GET("/shapes/:shapeType", h.sc.GetShapes)
		r.POST("/shapes", h.sc.CreateShape)
		ginLambda = ginadapter.New(r)

	}

	return ginLambda.ProxyWithContext(ctx, req)
}

func New(sc shapes.IShapesController) Handler {
	return Handler{
		sc: sc,
	}
}

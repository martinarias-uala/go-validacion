package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	shapes "github.com/martinarias-uala/go-validacion/internal/controller"
	"github.com/martinarias-uala/go-validacion/internal/handler"
	"github.com/martinarias-uala/go-validacion/internal/repository/dynamo"
)

func main() {
	d := dynamo.New()
	sc := shapes.New(d)
	h := handler.New(sc)

	lambda.Start(h.Handle)

}

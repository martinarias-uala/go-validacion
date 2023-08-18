package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/martinarias-uala/go-validacion/internal/handler"
	"github.com/martinarias-uala/go-validacion/internal/shapes"
)

func main() {
	sc := shapes.New()
	h := handler.New(&sc)

	lambda.Start(h.Handle)

}

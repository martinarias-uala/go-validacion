package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	shapes "github.com/martinarias-uala/go-validacion/internal/controller"
	"github.com/martinarias-uala/go-validacion/internal/handler"
	"github.com/martinarias-uala/go-validacion/internal/repository/dynamo"
	"github.com/martinarias-uala/go-validacion/internal/repository/s3"
)

func main() {
	s3 := s3.New()
	d := dynamo.New()
	sc := shapes.New(d, s3)
	h := handler.New(sc)

	lambda.Start(h.Handle)

}

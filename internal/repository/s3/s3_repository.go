package s3

/*

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/martinarias-uala/go-validacion/pkg/models"
)

type S3Client interface {
	PutObject(context.Context, *s3.PutObjectInput, ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

type S3R interface {
	PutObject(string, string, []models.Shape) models.ACTraceItem
}

type S3 struct {
	client S3Client
}

func New() *S3 {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := s3.NewFromConfig(cfg)
	return &S3{
		client: client,
	}
}

func (r *S3) PutObject(bucket, bucket_key string, loan []models.Shape) models.ACTraceItem {

	content, err := json.Marshal(loan)
	if err != nil {
		return util.NewACTraceItem("Error on parse object: " + err.Error())
	}

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(bucket_key),
		Body:   bytes.NewReader(content),
	}

	_, err = r.client.PutObject(context.TODO(), input)

	if err != nil {
		return util.NewACTraceItem("Error on put object: " + err.Error())
	}

	return models.ACTraceItem{}
}
*/

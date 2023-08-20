package s3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/martinarias-uala/go-validacion/pkg/models"
	"github.com/martinarias-uala/go-validacion/pkg/utils"
)

type S3Client interface {
	PutObject(context.Context, *s3.PutObjectInput, ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

type S3R interface {
	PutObject(shapes []models.ShapeData, shapeType string) error
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

func (r *S3) PutObject(shapes []models.ShapeData, shapeType string) error {

	content, err := json.Marshal(shapes)
	if err != nil {
		return err
	}

	currentTime := time.Now()
	formattedDate := currentTime.Format("2006-01-02")

	fileName := fmt.Sprintf("%s-%s-%v", shapeType, utils.GetAWSReqId(), formattedDate)
	bucketKey := fmt.Sprintf("SHAPES/%s.txt", fileName)

	input := &s3.PutObjectInput{
		Bucket: aws.String("uala-arg-labssupport-dev"),
		Key:    aws.String(bucketKey),
		Body:   bytes.NewReader(content),
	}

	_, err = r.client.PutObject(context.TODO(), input)

	if err != nil {
		return err
	}

	return nil
}

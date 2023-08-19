package s3

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/martinarias-uala/go-validacion/pkg/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("S3 Repository", func() {
	Context("Put Object path", putObjectPath)
})

func putObjectPath() {
	Context("Should Put Successfully", func() {

		shape := models.ShapeData{
			A: 23,
			B: 5,
			ShapeMetadata: models.ShapeMetadata{
				Type:      "TRIANGLE",
				ID:        "some-id",
				CreatedBy: "Leo Messi",
			},
		}
		c := MockS3Client{}
		out := s3.PutObjectOutput{}

		c.On("PutObject", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(&out, nil)
		r := S3{client: &c}

		err := r.PutObject(shape)
		It("Should not return error", func() {
			Ω(err).ToNot(HaveOccurred())
		})
	})

	Context("Should Fail on PutObject", func() {
		shape := models.ShapeData{
			A: 23,
			B: 5,
			ShapeMetadata: models.ShapeMetadata{
				Type:      "TRIANGLE",
				ID:        "some-id",
				CreatedBy: "Leo Messi",
			},
		}
		c := MockS3Client{}

		c.On("PutObject", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("some error"))
		r := S3{client: &c}

		err := r.PutObject(shape)
		It("Should throw error", func() {
			Ω(err).ToNot(BeNil())
		})
		It("Should have structured error", func() {
			Ω(err.Error()).To(ContainSubstring("some error"))
		})
	})
}

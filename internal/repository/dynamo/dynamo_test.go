package dynamo

import (
	"context"
	"errors"

	ddb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/martinarias-uala/go-validacion/pkg/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Dynamo Repository", func() {
	Context("Insert shape to dynamo happy path ", happyPath)
	Context("Insert shape to dynamo error path", errorPath)
})

func happyPath() {

	Context("Should Get shapes Successfully", func() {
		shape := models.ShapeData{
			A: 23,
			B: 5,
			ShapeMetadata: models.ShapeMetadata{
				Type:      "TRIANGLE",
				ID:        "some-id",
				CreatedBy: "Leo Messi",
			},
		}
		c := MockDynamoDBClient{}
		d := New()

		c.On("ExecuteStatement", context.TODO(), mock.Anything, mock.Anything).Return(&ddb.ExecuteStatementOutput{}, nil)
		d.client = &c

		data, err := d.GetShape(shape.Type, "")
		It("Should not return error", func() {
			Ω(err).To(BeNil())
		})

		It("Should return an empty array of Shapes", func() {
			Ω(data).To(Equal([]models.ShapeData{}))
		})
	})

	Context("Should Create a shape Successfully", func() {
		shape := models.ShapeData{
			A: 23,
			B: 5,
			ShapeMetadata: models.ShapeMetadata{
				Type:      "TRIANGLE",
				ID:        "some-id",
				CreatedBy: "Leo Messi",
			},
		}
		c := MockDynamoDBClient{}
		d := New()

		c.On("PutItem", context.TODO(), mock.Anything, mock.Anything).Return(&ddb.PutItemOutput{}, nil)
		d.client = &c

		err := d.CreateItem(shape)
		It("Should not return error", func() {
			Ω(err).To(BeNil())
		})
	})
}

func errorPath() {

	Context("Should fail on Get shapes", func() {
		shape := models.ShapeData{
			A: 23,
			B: 5,
			ShapeMetadata: models.ShapeMetadata{
				Type:      "TRIANGLE",
				ID:        "some-id",
				CreatedBy: "Leo Messi",
			},
		}
		c := MockDynamoDBClient{}
		d := New()

		c.On("ExecuteStatement", context.TODO(), mock.Anything, mock.Anything).Return(&ddb.ExecuteStatementOutput{}, errors.New("some error"))
		d.client = &c

		_, err := d.GetShape(shape.Type, "")
		It("Should return error", func() {
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(ContainSubstring("some error"))
		})
	})

	Context("Should fail on Create shape", func() {
		shape := models.ShapeData{
			A: 23,
			B: 5,
			ShapeMetadata: models.ShapeMetadata{
				Type:      "TRIANGLE",
				ID:        "some-id",
				CreatedBy: "Leo Messi",
			},
		}
		c := MockDynamoDBClient{}
		d := New()

		c.On("PutItem", context.TODO(), mock.Anything, mock.Anything).Return(&ddb.PutItemOutput{}, errors.New("some error"))
		d.client = &c

		err := d.CreateItem(shape)
		It("Should return error", func() {
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(ContainSubstring("some error"))
		})
	})
}

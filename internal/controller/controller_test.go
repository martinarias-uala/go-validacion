package shapes

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/martinarias-uala/go-validacion/internal/repository/dynamo"
	"github.com/martinarias-uala/go-validacion/internal/repository/s3"
	"github.com/martinarias-uala/go-validacion/pkg/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Shapes Controller", func() {
	Context("Get Shapes path", getShapesPath)
})

func getShapesPath() {
	Context("Should Get Shapes Successfully", func() {
		expectedShapes := []models.ShapeData{
			{
				A: 23,
				B: 5,
				ShapeMetadata: models.ShapeMetadata{
					Type:      "TRIANGLE",
					ID:        "some-id",
					CreatedBy: "Leo Messi",
				},
			},
		}
		dynamoMock := &dynamo.MockDynamoDBRepository{}
		s3Mock := &s3.MockS3Repository{}

		dynamoMock.On("GetShape", mock.Anything).Return(expectedShapes, nil)
		s3Mock.On("PutObject", mock.Anything).Return(nil)

		sc := New(dynamoMock, s3Mock)

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Params = append(c.Params, gin.Param{
			Key:   "shapeType",
			Value: "TRIANGLE",
		})
		sc.GetShapes(c)

		It("Should have status 200", func() {
			Ω(c.Writer.Status()).Should(Equal(200))
		})
		It("Should call s3.PutObject", func() {
			Ω(s3Mock.ExpectedCalls[0].Method).To(Equal("PutObject"))
		})
		It("Should call dynamo.GetShape", func() {
			Ω(dynamoMock.ExpectedCalls[0].Method).To(Equal("GetShape"))
		})
	})

	/*
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
				c := MockClient{}

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
	*/
}

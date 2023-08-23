package shapes

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/martinarias-uala/go-validacion/internal/repository/dynamo"
	"github.com/martinarias-uala/go-validacion/internal/repository/s3"
	"github.com/martinarias-uala/go-validacion/pkg/models"
	"github.com/martinarias-uala/go-validacion/pkg/service"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Shapes Controller", func() {
	Context("Get Shapes path", getShapesPath)
	Context("Create Shape path", createShapePath)
})

func getShapesPath() {
	Context("Should Get Shapes Successfully from DynamoDB", func() {

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
		httpClientMock := &service.MockHttp{}
		dynamoMock.On("GetShape", "TRIANGLE", "").Return(models.GetShapesResponse{
			ShapesData: expectedShapes,
		}, nil)

		s3Mock.On("PutObject", mock.Anything, mock.Anything).Return(nil)

		sc := New(dynamoMock, s3Mock, httpClientMock)

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
	Context("Should Put Shape on S3 Successfully", func() {

		ellipse := models.Ellipse{
			SemiMajorAxis: 10,
			SemiMinorAxis: 10,
		}
		shapeToPut := ellipse.ToGenericShape(models.ShapeMetadata{
			Type:      "ELLIPSE",
			ID:        "some-id",
			CreatedBy: "Leo Messi",
			Area:      ellipse.CalculateArea(),
		})
		expectedShapes := []models.ShapeData{
			shapeToPut,
		}

		dynamoMock := &dynamo.MockDynamoDBRepository{}
		s3Mock := &s3.MockS3Repository{}
		httpClientMock := &service.MockHttp{}

		dynamoMock.On("GetShape", mock.Anything, mock.Anything).Return(models.GetShapesResponse{
			ShapesData: expectedShapes,
		}, nil)

		s3Mock.On("PutObject", expectedShapes, "ELLIPSE").Return(nil)

		sc := New(dynamoMock, s3Mock, httpClientMock)

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Params = append(c.Params, gin.Param{
			Key:   "shapeType",
			Value: "ELLIPSE",
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
	Context("Should fail on Get Shapes from DynamoDB", func() {

		dynamoMock := &dynamo.MockDynamoDBRepository{}
		s3Mock := &s3.MockS3Repository{}
		httpClientMock := &service.MockHttp{}

		dynamoMock.On("GetShape", mock.Anything, mock.Anything).Return(models.GetShapesResponse{}, errors.New("Some error"))

		s3Mock.On("PutObject", mock.Anything, mock.Anything).Return(nil)

		sc := New(dynamoMock, s3Mock, httpClientMock)

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Params = append(c.Params, gin.Param{
			Key:   "shapeType",
			Value: "TRIANGLE",
		})
		sc.GetShapes(c)

		It("Should have status 500", func() {
			Ω(c.Writer.Status()).Should(Equal(500))
		})

		It("Should call dynamo.GetShape", func() {
			Ω(dynamoMock.ExpectedCalls[0].Method).To(Equal("GetShape"))
		})
	})
	Context("Should fail on Put Shape on S3", func() {

		ellipse := models.Ellipse{
			SemiMajorAxis: 10,
			SemiMinorAxis: 10,
		}
		shapeToPut := ellipse.ToGenericShape(models.ShapeMetadata{
			Type:      "RECTANGLE",
			ID:        "some-id",
			CreatedBy: "Leo Messi",
			Area:      ellipse.CalculateArea(),
		})
		expectedShapes := []models.ShapeData{
			shapeToPut,
		}

		dynamoMock := &dynamo.MockDynamoDBRepository{}
		s3Mock := &s3.MockS3Repository{}
		httpClientMock := &service.MockHttp{}

		dynamoMock.On("GetShape", mock.Anything, mock.Anything).Return(models.GetShapesResponse{
			ShapesData: expectedShapes,
		}, nil)

		s3Mock.On("PutObject", mock.Anything, mock.Anything).Return(errors.New("some-error"))

		sc := New(dynamoMock, s3Mock, httpClientMock)

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Params = append(c.Params, gin.Param{
			Key:   "shapeType",
			Value: "RECTANGLE",
		})
		sc.GetShapes(c)

		It("Should have status 500", func() {
			Ω(c.Writer.Status()).Should(Equal(500))
		})
		It("Should call s3.PutObject", func() {
			Ω(s3Mock.ExpectedCalls[0].Method).To(Equal("PutObject"))
		})
		It("Should call dynamo.GetShape", func() {
			Ω(dynamoMock.ExpectedCalls[0].Method).To(Equal("GetShape"))
		})
	})
}
func createShapePath() {
	Context("Should create item successfully", func() {

		stringReader := strings.NewReader(`{"data":{"email":"some-email@email.com"}}`)
		stringReadCloser := io.NopCloser(stringReader)
		resp := http.Response{
			StatusCode: http.StatusOK,
			Body:       stringReadCloser,
		}

		dynamoMock := &dynamo.MockDynamoDBRepository{}
		s3Mock := &s3.MockS3Repository{}
		httpClientMock := &service.MockHttp{}

		httpClientMock.On("Get", mock.Anything).Return(&resp, nil)
		dynamoMock.On("CreateItem", mock.Anything).Return(nil)
		sc := New(dynamoMock, s3Mock, httpClientMock)

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Params = append(c.Params, gin.Param{
			Key:   "shapeType",
			Value: "TRIANGLE",
		})
		c.Request, _ = http.NewRequest("GET", "?id=2&a=1&b=1", nil)

		sc.CreateShape(c)

		It("Should have status 201", func() {
			Ω(c.Writer.Status()).Should(Equal(201))
		})
		It("Should call dynamo.GetShape", func() {
			Ω(dynamoMock.ExpectedCalls[0].Method).To(Equal("CreateItem"))
		})
	})
	Context("Should fail creating item ", func() {

		stringReader := strings.NewReader(`{"data":{"email":"some-email@email.com"}}`)
		stringReadCloser := io.NopCloser(stringReader)
		resp := http.Response{
			StatusCode: http.StatusOK,
			Body:       stringReadCloser,
		}

		dynamoMock := &dynamo.MockDynamoDBRepository{}
		s3Mock := &s3.MockS3Repository{}
		httpClientMock := &service.MockHttp{}

		httpClientMock.On("Get", mock.Anything).Return(&resp, nil)
		dynamoMock.On("CreateItem", mock.Anything).Return(errors.New("some-error"))
		sc := New(dynamoMock, s3Mock, httpClientMock)

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Params = append(c.Params, gin.Param{
			Key:   "shapeType",
			Value: "TRIANGLE",
		})
		c.Request, _ = http.NewRequest("GET", "?id=2&a=1&b=1", nil)

		sc.CreateShape(c)

		It("Should have status 500", func() {
			Ω(c.Writer.Status()).Should(Equal(500))
		})
		It("Should call dynamo.GetShape", func() {
			Ω(dynamoMock.ExpectedCalls[0].Method).To(Equal("CreateItem"))
		})
	})

}

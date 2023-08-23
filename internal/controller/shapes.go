package shapes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/martinarias-uala/go-validacion/internal/repository/dynamo"
	"github.com/martinarias-uala/go-validacion/internal/repository/s3"
	"github.com/martinarias-uala/go-validacion/pkg/models"
	"github.com/martinarias-uala/go-validacion/pkg/service"
	"github.com/martinarias-uala/go-validacion/pkg/utils"
)

type IShapesController interface {
	GetShapes(*gin.Context)
	CreateShape(*gin.Context)
}

type ShapesController struct {
	r   dynamo.DynamoRepository
	s3r s3.S3R
	h   service.HTTPClient
}

func New(r dynamo.DynamoRepository, s3r s3.S3R, h service.HTTPClient) *ShapesController {
	return &ShapesController{
		r:   r,
		s3r: s3r,
		h:   h,
	}
}

func (sc *ShapesController) CreateShape(c *gin.Context) {
	var responseData models.ResponseData

	shapeType := c.Param("shapeType")
	id := c.Query("id")
	aStr := c.Query("a")
	bStr := c.Query("b")

	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	requestUrl := fmt.Sprintf("https://reqres.in/api/users/%s", id)
	response, err := sc.h.Get(requestUrl)

	if err != nil {
		c.JSON(response.StatusCode, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = json.Unmarshal(body, &responseData)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = sc.r.CreateItem(models.ShapeData{
		A: a,
		B: b,
		ShapeMetadata: models.ShapeMetadata{
			ID:        utils.GetUUID(),
			Type:      shapeType,
			CreatedBy: responseData.Data.Email,
		},
	})
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (sc *ShapesController) GetShapes(c *gin.Context) {
	var shapesToPut []models.ShapeData

	shapeType := c.Param("shapeType")
	nextToken := c.Query("page")

	shapesResponse, err := sc.r.GetShape(shapeType, nextToken)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})

	}

	shapes := shapesResponse.ShapesData
	pageToken := shapesResponse.PageToken

	var wg sync.WaitGroup
	shapesCh := make(chan models.ShapeData)

	for _, v := range shapes {
		wg.Add(1)
		go func(shapeData models.ShapeData) {
			defer wg.Done()

			var shape models.ShapeData
			switch shapeData.Type {
			case "RECTANGLE":
				rectangle := models.Rectangle{
					Length: shapeData.A,
					Width:  shapeData.B,
				}
				shape = rectangle.ToGenericShape(models.ShapeMetadata{
					ID:        shapeData.ID,
					CreatedBy: shapeData.CreatedBy,
					Type:      shapeData.Type,
				})
				shape.Area = rectangle.CalculateArea()

			case "ELLIPSE":
				ellipse := models.Ellipse{
					SemiMajorAxis: shapeData.A,
					SemiMinorAxis: shapeData.B,
				}
				shape = ellipse.ToGenericShape(models.ShapeMetadata{
					ID:        shapeData.ID,
					CreatedBy: shapeData.CreatedBy,
					Type:      shapeData.Type,
				})
				shape.Area = ellipse.CalculateArea()

			case "TRIANGLE":
				triangle := models.Triangle{
					Base:   shapeData.A,
					Height: shapeData.B,
				}
				shape = triangle.ToGenericShape(models.ShapeMetadata{
					ID:        shapeData.ID,
					CreatedBy: shapeData.CreatedBy,
					Type:      shapeData.Type,
				})
				shape.Area = triangle.CalculateArea()
			}
			shapesCh <- shape
		}(v)
	}

	go func() {
		wg.Wait()
		close(shapesCh)
	}()

	for shapeData := range shapesCh {
		shapesToPut = append(shapesToPut, shapeData)
	}

	err = sc.s3r.PutObject(shapesToPut, shapeType)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})

	}

	c.JSON(200, gin.H{
		"data":       shapesToPut,
		"page_token": pageToken,
	})
}

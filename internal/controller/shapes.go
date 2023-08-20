package shapes

import (
	"github.com/gin-gonic/gin"
	"github.com/martinarias-uala/go-validacion/internal/repository/dynamo"
	"github.com/martinarias-uala/go-validacion/internal/repository/s3"
	"github.com/martinarias-uala/go-validacion/pkg/models"
)

type IShapesController interface {
	GetShapes(*gin.Context)
	CreateShape(*gin.Context)
}

type ShapesController struct {
	r   dynamo.DynamoRepository
	s3r s3.S3R
}

func New(r dynamo.DynamoRepository, s3r s3.S3R) *ShapesController {
	return &ShapesController{
		r:   r,
		s3r: s3r,
	}
}

func (sc *ShapesController) CreateShape(c *gin.Context) {
	/* shapeType := c.Param("shapeType")
	switch shapeType {
	case "RECTANGLE":
		c.JSON(200, models.Rectangle{})
	case "ELLIPSE":
		c.JSON(200, models.Ellipse{})
	case "TRIANGLE":
		c.JSON(200, models.Triangle{})
	}
	err := c.BindJSON(&newShape)

	if err != nil {
		return
	}

	id := utils.GetUUID()
	c.JSON(http.StatusCreated, newShape) */

}

func (sc *ShapesController) GetShapes(c *gin.Context) {
	var shapesToPut []models.ShapeData

	shapeType := c.Param("shapeType")
	shapes, err := sc.r.GetShape(shapeType)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})

	}

	for _, v := range shapes {
		switch v.Type {
		case "RECTANGLE":
			rectangle := models.Rectangle{
				Length: v.A,
				Width:  v.B,
			}
			shapesToPut = append(shapesToPut, rectangle.ToGenericShape(models.ShapeMetadata{
				ID:        v.ID,
				CreatedBy: v.CreatedBy,
				Type:      v.Type,
				Area:      rectangle.CalculateArea(),
			}))

		case "ELLIPSE":
			ellipse := models.Ellipse{
				SemiMajorAxis: v.A,
				SemiMinorAxis: v.B,
			}
			shapesToPut = append(shapesToPut, ellipse.ToGenericShape(models.ShapeMetadata{
				ID:        v.ID,
				CreatedBy: v.CreatedBy,
				Type:      v.Type,
				Area:      ellipse.CalculateArea(),
			}))

		case "TRIANGLE":
			triangle := models.Triangle{
				Base:   v.A,
				Height: v.B,
			}
			shapesToPut = append(shapesToPut, triangle.ToGenericShape(models.ShapeMetadata{
				ID:        v.ID,
				CreatedBy: v.CreatedBy,
				Type:      v.Type,
				Area:      triangle.CalculateArea(),
			}))

		}
	}

	err = sc.s3r.PutObject(shapesToPut, shapeType)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})

	}

	c.JSON(200, shapes)

	/* limit := 10
	if c.Query("limit") != "" {
		newLimit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			limit = 10
		} else {
			limit = newLimit
		}
	}
	if limit > 50 {
		limit = 50
	}
	shapes := make([]models.Shape, limit)

	c.JSON(200, shapes) */
}

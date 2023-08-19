package shapes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/martinarias-uala/go-validacion/internal/repository/dynamo"
	"github.com/martinarias-uala/go-validacion/internal/repository/s3"
	"github.com/martinarias-uala/go-validacion/pkg/models"
)

type IShapesController interface {
	GetShape(*gin.Context)
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
	/*
		newShape := models.Rectangle{}
		err := c.BindJSON(&newShape)

		if err != nil {
			return
		}

		id := utils.GetUUID()
		newShape. = id
		c.JSON(http.StatusAccepted, newShape)
	*/
}

func (sc *ShapesController) GetShape(c *gin.Context) {
	/*
		shapeType := c.Param("shapeType")

		switch shapeType {
		case "RECTANGLE":
			c.JSON(200, models.Rectangle{})
		case "ELLIPSE":
			c.JSON(200, models.Ellipse{})
		case "TRIANGLE":
			c.JSON(200, models.Triangle{})
		} */

}

func (sc *ShapesController) GetShapes(c *gin.Context) {
	shapeType := c.Param("shapeType")

	shapes, _ := sc.r.GetShape(shapeType)

	for _, v := range shapes {
		switch v.Type {
		case "RECTANGLE":
			shape := models.Rectangle{
				Length: v.A,
				Width:  v.B,
			}

			sc.s3r.PutObject(shape.ToDynamoItem(models.ShapeMetadata{
				ID:        v.ID,
				CreatedBy: v.CreatedBy,
				Type:      v.Type,
				Area:      shape.CalculateArea(),
			}))

		case "ELLIPSE":
			shape := models.Ellipse{
				SemiMajorAxis: v.A,
				SemiMinorAxis: v.B,
			}
			log.Println(v)
			sc.s3r.PutObject(shape.ToDynamoItem(models.ShapeMetadata{
				ID:        v.ID,
				CreatedBy: v.CreatedBy,
				Type:      v.Type,
				Area:      shape.CalculateArea(),
			}))

		case "TRIANGLE":
			shape := models.Triangle{
				Base:   v.A,
				Height: v.B,
			}
			sc.s3r.PutObject(shape.ToDynamoItem(models.ShapeMetadata{
				ID:        v.ID,
				CreatedBy: v.CreatedBy,
				Type:      v.Type,
				Area:      shape.CalculateArea(),
			}))

		}
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

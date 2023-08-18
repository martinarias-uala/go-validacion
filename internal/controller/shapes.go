package shapes

import (
	"github.com/gin-gonic/gin"
	"github.com/martinarias-uala/go-validacion/internal/repository/dynamo"
)

type IShapesController interface {
	GetShape(*gin.Context)
	GetShapes(*gin.Context)
	CreateShape(*gin.Context)
}

type ShapesController struct {
	r dynamo.DynamoRepository
}

func New(r dynamo.DynamoRepository) *ShapesController {
	return &ShapesController{
		r: r,
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

	/* switch shapeType {
	case "RECTANGLE":
		c.JSON(200, models.Rectangle{})
	case "ELLIPSE":
		c.JSON(200, models.Ellipse{})
	case "TRIANGLE":
		c.JSON(200, models.Triangle{})
	} */

	shapes, _ := sc.r.GetShape(shapeType)

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

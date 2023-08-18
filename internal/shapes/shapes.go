package shapes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martinarias-uala/go-validacion/pkg/models"
	"github.com/martinarias-uala/go-validacion/pkg/utils"
)

type IShapesController interface {
	GetShape(*gin.Context)
	GetShapes(*gin.Context)
	CreateShape(*gin.Context)
}

type ShapesController struct {
}

func New() ShapesController {
	return ShapesController{}
}

func (sc *ShapesController) CreateShape(c *gin.Context) {

	newShape := models.Shape{}
	err := c.BindJSON(&newShape)

	if err != nil {
		return
	}

	id := utils.GetUUID()
	newShape.ID = id
	c.JSON(http.StatusAccepted, newShape)
}

func (sc *ShapesController) GetShape(c *gin.Context) {

	shapeType := c.Param("shapeType")

	c.JSON(200, models.Shape{
		Type: shapeType,
	})
}

func (sc *ShapesController) GetShapes(c *gin.Context) {

	limit := 10
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

	c.JSON(200, shapes)
}

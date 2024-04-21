package controller

import (
	"Gine2/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Controller struct {
	crud models.VideoModelInterface
}

type ControllerInterFace interface {
	Save(request *gin.Context) models.Video
	FindAll() []models.Video
	FindOne(request *gin.Context) (models.Video, error)
}

func NewVideoController(crud models.VideoModelInterface) ControllerInterFace {
	return &Controller{
		crud: crud,
	}
}

func (c *Controller) FindAll() []models.Video {
	return c.crud.FindAll()
}

var pk int = 0

func (c *Controller) Save(request *gin.Context) models.Video {
	var video models.Video
	pk += 1
	video.ID = pk
	request.BindJSON(&video)
	fmt.Println("Here is the video : ", video)
	c.crud.Save(video)
	return video
}
func (c *Controller) FindOne(request *gin.Context) (models.Video, error) {
	var param = request.Param("pk")
	pk, err := strconv.Atoi(param)
	if err == nil {
		result, err2 := c.crud.FindOne(pk)
		if err2 == nil {
			return result, nil
		}
		return result, err2
	}
	return models.Video{}, errors.New("bad format")

}

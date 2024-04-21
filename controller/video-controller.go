package controller

import (
	"Gine2/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	crud models.VideoModelInterface
}

type ControllerInterFace interface {
	Save(request *gin.Context) models.Video
	FindAll() []models.Video
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

package main

import (
	"Gine2/controller"
	"Gine2/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	videoModel      models.VideoModelInterface     = models.NewVideoModel()
	videoController controller.ControllerInterFace = controller.NewVideoController(videoModel)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(request *gin.Context) {
		request.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(request *gin.Context) {
		request.JSON(http.StatusCreated, videoController.Save(request))
	})

	server.GET("/videos/:pk", func(request *gin.Context) {
		query, err := videoController.FindOne(request)
		if err == nil {
			request.JSON(http.StatusOK, query)
		} else {
			request.JSON(http.StatusBadRequest, nil)
		}

	})
	server.Run(":8000")

}

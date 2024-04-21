package models

import (
	"errors"
	"fmt"
)

type Video struct {
	ID          int    `json:"ID,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"URL"`
}

type VideoModel struct {
	VideoList []Video
}

type VideoModelInterface interface {
	Save(Video) Video
	FindAll() []Video
	FindOne(pk int) (Video, error)
}

func NewVideoModel() VideoModelInterface {
	return &VideoModel{}
}

func (videoModel *VideoModel) FindAll() []Video {

	return videoModel.VideoList
}

func (videoModel *VideoModel) Save(video Video) Video {
	videoModel.VideoList = append(videoModel.VideoList, video)
	return video
}

func (videoModel *VideoModel) FindOne(pk int) (Video, error) {
	for _, item := range videoModel.VideoList {
		fmt.Println("www")
		if item.ID == pk {
			fmt.Println("Here is :", item)
			return item, nil
		}
	}
	return Video{}, errors.New("not found")
}

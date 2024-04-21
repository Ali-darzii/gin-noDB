package models

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

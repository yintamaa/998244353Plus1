package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yintamaa/998244353Plus1/repository"
	"github.com/yintamaa/998244353Plus1/service"
)

type FeedResponse struct {
	Response
	VideoList []*repository.Vedio `json:"video_list,omitempty"`
	NextTime  int64               `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	data, err := service.QueryVedioInfoFlow()
	if err != nil {

	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: data,
		NextTime:  time.Now().Unix(),
	})
}

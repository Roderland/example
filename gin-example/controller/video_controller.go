package controller

import (
	"gin-example/model"
	"github.com/gin-gonic/gin"
	"sync"
)

type VideoController struct {
	videos      []model.Video
	idGenerator VideoIdGenerator
}

type VideoIdGenerator struct {
	counter int
	mtx     sync.Mutex
}

func (g *VideoIdGenerator) getNextId() int {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	g.counter++
	return g.counter
}

func (g *VideoIdGenerator) backId() {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	g.counter--
}

func NewVideoController() *VideoController {
	return &VideoController{
		videos: []model.Video{
			{Id: 1, Title: "test1", Description: "test1"},
			{Id: 2, Title: "test2", Description: "test2"},
		},
		idGenerator: VideoIdGenerator{
			counter: 2,
		},
	}
}

func (vc *VideoController) GetAll(context *gin.Context) {
	context.JSON(200, vc.videos)
}

func (vc *VideoController) Create(context *gin.Context) {
	video := model.Video{Id: vc.idGenerator.getNextId()}
	if err := context.ShouldBind(&video); err != nil {
		vc.idGenerator.backId()
		context.String(400, "bad request %v", err)
		return
	}
	vc.videos = append(vc.videos, video)
	context.String(200, "success")
}

func (vc *VideoController) Update(context *gin.Context) {
	var target model.Video
	if err := context.ShouldBindUri(&target); err != nil {
		context.String(400, "bad request %v", err)
		return
	}
	if err := context.ShouldBind(&target); err != nil {
		context.String(400, "bad request %v", err)
		return
	}
	for i, video := range vc.videos {
		if video.Id == target.Id {
			vc.videos[i] = target
			context.String(200, "update success")
			return
		}
	}
	context.String(400, "cannot find video-%d", target.Id)
}

func (vc *VideoController) Delete(context *gin.Context) {
	var target model.Video
	if err := context.ShouldBindUri(&target); err != nil {
		context.String(400, "bad request %v", err)
		return
	}
	for i, video := range vc.videos {
		if video.Id == target.Id {
			vc.videos = append(vc.videos[:i], vc.videos[i+1:]...)
			context.String(200, "delete success")
			return
		}
	}
	context.String(400, "cannot find video-%d", target.Id)
}

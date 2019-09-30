package router

import (
	"github.com/gin-gonic/gin"
	"test/models/engine"
)

type Basic struct {}

type BasicSimpleView struct {
	Containers int `json:"containers"`
	ContainersRunning int `json:"containers_running"`
	ContainersPaused int `json:"containers_paused"`
	ContainersStopped int `json:"containers_stopped"`
	Images int `json:"images"`
}

func (basic *Basic) GetBasicInfo(c *gin.Context) {
	cli := engine.GetDockerEngine()
	info, err := cli.Cli.Info(cli.Ctx)
	if err != nil {
		Result(c, STATUS_FAILED, nil, err)
	}
	Result(c, STATUS_OK, BasicSimpleView{
		Containers:        info.Containers,
		ContainersRunning: info.ContainersRunning,
		ContainersPaused:  info.ContainersPaused,
		ContainersStopped: info.ContainersStopped,
		Images:            info.Images,
	}, nil)
}

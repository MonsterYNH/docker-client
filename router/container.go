package router

import (
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"test/models/engine"
)

type Container struct {}

func (container *Container) GetContainerList(c *gin.Context) {
	cli := engine.GetDockerEngine()
	containers, err := cli.Cli.ContainerList(cli.Ctx, types.ContainerListOptions{
		All:     true,
	})
	if err != nil {
		Result(c, STATUS_FAILED, nil, err)
	}
	Result(c, STATUS_OK, containers, nil)
}

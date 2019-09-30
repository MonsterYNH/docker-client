package engine

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
)

var engine *DockerEngine

func init() {
	var err error
	engine, err = CreateDockerEngine()
	if err != nil {
		panic(fmt.Sprintf("ERROR: connect docker engine failed, error: %s", err))
	}
}

type DockerEngine struct {
	Ctx context.Context
	Cli *client.Client
}

func CreateDockerEngine() (*DockerEngine, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &DockerEngine{
		Ctx: ctx,
		Cli: cli,
	}, nil
}

func GetDockerEngine() *DockerEngine {
	return engine
}

package nerd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"

	"code.cloudfoundry.org/garden"
)

type Garden struct {
	Containerd *containerd.Client
	Logger     *log.Logger

	containers map[string]*Container
}

func New(containerdClient *containerd.Client, logger *log.Logger) *Garden {
	return &Garden{
		Containerd: containerdClient,
		Logger:     logger,
		containers: map[string]*Container{},
	}
}

func (g *Garden) Ping() error {
	return nil
}

func (g *Garden) Capacity() (garden.Capacity, error) {
	return garden.Capacity{}, nil
}

func (g *Garden) Destroy(handle string) error {
	return nil
}

func (g *Garden) Containers(garden.Properties) ([]garden.Container, error) {
	var containers []garden.Container
	for _, container := range g.containers {
		containers = append(containers, container)
	}
	return containers, nil
}

func (g *Garden) BulkInfo(handles []string) (map[string]garden.ContainerInfoEntry, error) {
	return nil, nil
}

func (g *Garden) BulkMetrics(handles []string) (map[string]garden.ContainerMetricsEntry, error) {
	return nil, nil
}

func (g *Garden) Lookup(handle string) (garden.Container, error) {
	ctr, ok := g.containers[handle]
	if !ok {
		return nil, fmt.Errorf("no container with handle %s found", handle)
	}
	return ctr, nil
}

func (g *Garden) Start() error {
	return nil
}

func (g *Garden) Stop() {
}

func (g *Garden) GraceTime(garden.Container) time.Duration {
	return time.Minute
}

func namespace() context.Context {
	return namespaces.WithNamespace(context.Background(), "garden")
}

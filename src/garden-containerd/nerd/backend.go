package nerd

import (
	"log"
	"time"

	"github.com/containerd/containerd"

	"code.cloudfoundry.org/garden"
)

type Garden struct {
	Containerd *containerd.Client
	Logger     *log.Logger
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
	return nil, nil
}

func (g *Garden) BulkInfo(handles []string) (map[string]garden.ContainerInfoEntry, error) {
	return nil, nil
}

func (g *Garden) BulkMetrics(handles []string) (map[string]garden.ContainerMetricsEntry, error) {
	return nil, nil
}

func (g *Garden) Lookup(handle string) (garden.Container, error) {
	return nil, nil
}

func (g *Garden) Start() error {
	return nil
}

func (g *Garden) Stop() {
}

func (g *Garden) GraceTime(garden.Container) time.Duration {
	return time.Minute
}

package nerd

import (
	"io"
	"time"

	"code.cloudfoundry.org/garden"

	"github.com/containerd/containerd"
)

type Container struct {
	ctr    containerd.Container
	handle string
}

func (c *Container) Handle() string {
	return c.handle
}

func (c *Container) Stop(kill bool) error {
	return nil
}

func (c *Container) Info() (garden.ContainerInfo, error) {
	return garden.ContainerInfo{}, nil
}

func (c *Container) StreamIn(spec garden.StreamInSpec) error {
	return nil
}

func (c *Container) StreamOut(spec garden.StreamOutSpec) (io.ReadCloser, error) {
	return nil, nil
}

func (c *Container) CurrentBandwidthLimits() (garden.BandwidthLimits, error) {
	return garden.BandwidthLimits{}, nil
}

func (c *Container) CurrentCPULimits() (garden.CPULimits, error) {
	return garden.CPULimits{}, nil
}

func (c *Container) CurrentDiskLimits() (garden.DiskLimits, error) {
	return garden.DiskLimits{}, nil
}

func (c *Container) CurrentMemoryLimits() (garden.MemoryLimits, error) {
	return garden.MemoryLimits{}, nil
}

func (c *Container) NetIn(hostPort, containerPort uint32) (uint32, uint32, error) {
	return 0, 0, nil
}

func (c *Container) NetOut(netOutRule garden.NetOutRule) error {
	return nil
}

func (c *Container) BulkNetOut(netOutRules []garden.NetOutRule) error {
	return nil
}

func (c *Container) Attach(processID string, io garden.ProcessIO) (garden.Process, error) {
	return nil, nil
}

func (c *Container) Metrics() (garden.Metrics, error) {
	return garden.Metrics{}, nil
}

func (c *Container) SetGraceTime(graceTime time.Duration) error {
	return nil
}

func (c *Container) Properties() (garden.Properties, error) {
	return nil, nil
}

func (c *Container) Property(name string) (string, error) {
	return "", nil
}

func (c *Container) SetProperty(name string, value string) error {
	return nil
}

func (c *Container) RemoveProperty(name string) error {
	return nil
}

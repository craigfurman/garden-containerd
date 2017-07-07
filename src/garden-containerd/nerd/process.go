package nerd

import (
	"fmt"

	"code.cloudfoundry.org/garden"
	"github.com/containerd/containerd"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/pborman/uuid"
)

func (c *Container) Run(spec garden.ProcessSpec, pio garden.ProcessIO) (garden.Process, error) {
	if spec.ID == "" {
		spec.ID = uuid.New()
	}
	if spec.Dir == "" {
		spec.Dir = "/"
	}

	ns := namespace()
	taskIO := containerd.NewIO(pio.Stdin, pio.Stdout, pio.Stderr)
	task, err := c.ctr.NewTask(ns, taskIO)
	if err != nil {
		return nil, fmt.Errorf("creating task: %s", err)
	}

	ociSpec := &specs.Process{
		Terminal: false,
		Args:     append([]string{spec.Path}, spec.Args...),
		Cwd:      spec.Dir,
	}

	proc, err := task.Exec(ns, spec.ID, ociSpec, taskIO)
	if err != nil {
		return nil, fmt.Errorf("task exec: %s", err)
	}

	if err := proc.Start(ns); err != nil {
		return nil, fmt.Errorf("starting process: %s", err)
	}

	return &process{proc: proc, id: spec.ID}, nil
}

type process struct {
	proc containerd.Process
	id   string
}

func (p *process) ID() string {
	return p.id
}

func (p *process) Wait() (int, error) {
	ns := namespace()
	defer p.proc.Delete(ns)
	exitStatus, err := p.proc.Wait(ns)
	return int(exitStatus), err
}

func (p *process) SetTTY(garden.TTYSpec) error {
	return nil
}

func (p *process) Signal(garden.Signal) error {
	return nil
}

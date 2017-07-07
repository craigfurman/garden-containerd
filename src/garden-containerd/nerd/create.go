package nerd

import (
	"code.cloudfoundry.org/garden"
	"github.com/containerd/containerd"
	"github.com/pborman/uuid"
)

func (g *Garden) Create(spec garden.ContainerSpec) (garden.Container, error) {
	if spec.Handle == "" {
		spec.Handle = uuid.New()
	}

	ns := namespace()
	image, err := g.Containerd.Pull(ns, spec.Image.URI, containerd.WithPullUnpack)
	if err != nil {
		g.Logger.Println(err)
		return nil, err
	}

	ociSpec, err := containerd.GenerateSpec(
		containerd.WithImageConfig(ns, image),
		containerd.WithProcessArgs("sleep", "3600"),
	)

	if err != nil {
		g.Logger.Println(err)
		return nil, err
	}

	ctr, err := g.Containerd.NewContainer(
		ns, spec.Handle,
		containerd.WithSpec(ociSpec),
		containerd.WithNewRootFS(spec.Handle, image),
	)
	if err != nil {
		g.Logger.Println(err)
		return nil, err
	}
	gdnCtr := &Container{ctr: ctr, handle: spec.Handle}

	g.containers[spec.Handle] = gdnCtr

	return gdnCtr, nil
}

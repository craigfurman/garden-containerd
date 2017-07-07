package nerd

import (
	"context"

	"code.cloudfoundry.org/garden"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
	"github.com/pborman/uuid"
)

func (g *Garden) Create(spec garden.ContainerSpec) (garden.Container, error) {
	if spec.Handle == "" {
		spec.Handle = uuid.New()
	}

	ctx := context.Background()
	ns := namespaces.WithNamespace(ctx, "garden")

	image, err := g.Containerd.Pull(ns, spec.Image.URI, containerd.WithPullUnpack)
	if err != nil {
		g.Logger.Println(err)
		return nil, err
	}

	ociSpec, err := containerd.GenerateSpec(
		containerd.WithImageConfig(ns, image),
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

	return &Container{ctr: ctr}, nil
}

package engine

import (
	"errors"
	"io"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/pivotal-golang/lager"
)

var ErrBuildNotFound = errors.New("build not found")

//go:generate counterfeiter . Engine

type Engine interface {
	Name() string

	CreateBuild(db.Build, atc.BuildPlan) (Build, error)
	LookupBuild(db.Build) (Build, error)
}

//go:generate counterfeiter . Build

type Build interface {
	Metadata() string

	Abort() error
	Hijack(atc.HijackProcessSpec, HijackProcessIO) (HijackedProcess, error)
	Resume(lager.Logger)
}

type HijackProcessIO struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

//go:generate counterfeiter . HijackedProcess

type HijackedProcess interface {
	Wait() (int, error)
	SetTTY(atc.HijackTTYSpec) error
}

type Engines []Engine

func (engines Engines) Lookup(name string) (Engine, bool) {
	for _, e := range engines {
		if e.Name() == name {
			return e, true
		}
	}

	return nil, false
}
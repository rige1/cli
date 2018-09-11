package types

import (
	"context"
	"io"

	registryclient "github.com/docker/cli/cli/registry/client"
	"github.com/docker/docker/api/types"
	ver "github.com/hashicorp/go-version"
)

const (
	// CommunityEngineImage is the repo name for the community engine
	CommunityEngineImage = "engine-community"

	// EnterpriseEngineImage is the repo name for the enterprise engine
	EnterpriseEngineImage = "engine-enterprise"
)

// ContainerizedClient can be used to manage the lifecycle of
// dockerd running as a container on containerd.
type ContainerizedClient interface {
	Close() error
	ActivateEngine(ctx context.Context,
		opts EngineInitOptions,
		out OutStream,
		authConfig *types.AuthConfig,
		healthfn func(context.Context) error) error
	InitEngine(ctx context.Context,
		opts EngineInitOptions,
		out OutStream,
		authConfig *types.AuthConfig,
		healthfn func(context.Context) error) error
	DoUpdate(ctx context.Context,
		opts EngineInitOptions,
		out OutStream,
		authConfig *types.AuthConfig,
		healthfn func(context.Context) error) error
	GetEngineVersions(ctx context.Context, registryClient registryclient.RegistryClient, currentVersion, imageName string) (AvailableVersions, error)
	GetCurrentEngineVersion(ctx context.Context) (EngineInitOptions, error)
	RemoveEngine(ctx context.Context) error
}

// EngineInitOptions contains the configuration settings
// use during initialization of a containerized docker engine
type EngineInitOptions struct {
	RegistryPrefix string
	EngineImage    string
	EngineVersion  string
	ConfigFile     string
	Scope          string
}

// AvailableVersions groups the available versions which were discovered
type AvailableVersions struct {
	Downgrades []DockerVersion
	Patches    []DockerVersion
	Upgrades   []DockerVersion
}

// DockerVersion wraps a semantic version to retain the original tag
// since the docker date based versions don't strictly follow semantic
// versioning (leading zeros, etc.)
type DockerVersion struct {
	ver.Version
	Tag string
}

// Update stores available updates for rendering in a table
type Update struct {
	Type    string
	Version string
	Notes   string
}

// OutStream is an output stream used to write normal program output.
type OutStream interface {
	io.Writer
	FD() uintptr
	IsTerminal() bool
}

package manager

import (
	"encoding/json"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	pluginNameRe = regexp.MustCompile("^[a-z][a-z0-9]*$")
)

// Plugin represents a potential plugin with all it's metadata.
type Plugin struct {
	Metadata

	Name string
	Path string

	// Err is non-nil if the plugin failed one of the candidate tests.
	Err error `json:",omitempty"`

	// ShadowedPaths contains the paths of any other plugins which this plugin takes precedence over.
	ShadowedPaths []string `json:",omitempty"`
}

// newPlugin determines if the given candidate is valid and returns a
// Plugin.  If the candidate fails one of the tests then `Plugin.Err`
// is set, but the `Plugin` is still returned with no error. An error
// is only returned due to a non-recoverable error.
func newPlugin(c Candidate, rootcmd *cobra.Command) (Plugin, error) {
	path := c.Path()
	if path == "" {
		return Plugin{}, errors.New("plugin candidate path cannot be empty")
	}

	// The candidate listing process should have skipped anything
	// which would fail here, so there are all real errors.
	fullname := filepath.Base(path)
	if fullname == "." {
		return Plugin{}, errors.Errorf("unable to determine basename of plugin candidate %q", path)
	}
	if runtime.GOOS == "windows" {
		exe := ".exe"
		if !strings.HasSuffix(fullname, exe) {
			return Plugin{}, errors.Errorf("plugin candidate %q lacks required %q suffix", path, exe)
		}
		fullname = strings.TrimSuffix(fullname, exe)
	}
	if !strings.HasPrefix(fullname, NamePrefix) {
		return Plugin{}, errors.Errorf("plugin candidate %q does not have %q prefix", path, NamePrefix)
	}

	p := Plugin{
		Name: strings.TrimPrefix(fullname, NamePrefix),
		Path: path,
	}

	// Now apply the candidate tests, so these update p.Err.
	if !pluginNameRe.MatchString(p.Name) {
		p.Err = errors.Errorf("plugin candidate %q did not match %q", p.Name, pluginNameRe.String())
		return p, nil
	}

	if rootcmd != nil {
		for _, cmd := range rootcmd.Commands() {
			if cmd.Name() == p.Name {
				p.Err = errors.Errorf("plugin %q duplicates builtin command", p.Name)
				return p, nil
			}
			if cmd.HasAlias(p.Name) {
				p.Err = errors.Errorf("plugin %q duplicates an alias of builtin command %q", p.Name, cmd.Name())
				return p, nil
			}
		}
	}

	// We are supposed to check for relevant execute permissions here. Instead we rely on an attempt to execute.
	meta, err := c.Metadata()
	if err != nil {
		p.Err = errors.Wrap(err, "failed to fetch metadata")
		return p, nil
	}

	if err := json.Unmarshal(meta, &p.Metadata); err != nil {
		p.Err = errors.Wrap(err, "invalid metadata")
		return p, nil
	}

	if p.Metadata.SchemaVersion != "0.1.0" {
		p.Err = errors.Errorf("plugin SchemaVersion %q is not valid, must be 0.1.0", p.Metadata.SchemaVersion)
		return p, nil
	}
	if p.Metadata.Vendor == "" {
		p.Err = errors.Errorf("plugin metadata does not define a vendor")
		return p, nil
	}
	return p, nil
}

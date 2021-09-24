package sidekiq

import (
	"fmt"
	"path/filepath"

	"github.com/paketo-buildpacks/packit"
)

type SidekiqExecutable struct{}

func NewSidekiqExecutable() SidekiqExecutable {
	return SidekiqExecutable{}
}

func (p SidekiqExecutable) Find(context packit.BuildContext) (string, error) {
	gemfileParser := NewGemfileParser()
	hasSidekiqPro, err := gemfileParser.Parse(filepath.Join(context.WorkingDir, "Gemfile"), "sidekiq-pro")

	// realistically it should never get this far, as if the Gemfile doesn't exist the detect step will fail
	if err != nil {
		return "", fmt.Errorf("failed to parse Gemfile: %w", err)
	}

	exec := "sidekiq"
	if hasSidekiqPro {
		exec = "sidekiqswarm"
	}

	return exec, nil
}

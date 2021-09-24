package main

import (
	"os"

	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/scribe"
	"github.com/paketo-buildpacks/sidekiq"
)

func main() {
	parser := sidekiq.NewGemfileParser()
	logger := scribe.NewLogger(os.Stdout)

	packit.Run(
		sidekiq.Detect(parser),
		sidekiq.Build(logger),
	)
}

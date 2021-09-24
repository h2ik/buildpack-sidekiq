package sidekiq

import (
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/scribe"
)

func Build(logger scribe.Logger) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		logger.Title("%s %s", context.BuildpackInfo.Name, context.BuildpackInfo.Version)

		command := "bundle exec sidekiq ${args}"

		logger.Process("Assigning launch processes")
		logger.Subprocess("sidekiq: %s", command)
		logger.Break()

		return packit.BuildResult{
			Launch: packit.LaunchMetadata{
				Processes: []packit.Process{
					{
						Type:    "sidekiq",
						Command: command,
					},
				},
			},
		}, nil
	}
}

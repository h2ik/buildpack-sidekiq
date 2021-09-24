package sidekiq

import (
	"fmt"

	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/scribe"
)

func Build(logger scribe.Logger) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		logger.Title("%s %s", context.BuildpackInfo.Name, context.BuildpackInfo.Version)

		sidekiqExec := NewSidekiqExecutable()
		exec, err := sidekiqExec.Find(context)
		if err != nil {
			return packit.BuildResult{}, fmt.Errorf("failed to find sidekiq executable: %w", err)
		}

		command := fmt.Sprintf("bundle exec %s -t ${timeout:-60} -c ${threads:-5} ${additional_args}", exec)

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

# Sidekiq Cloud Native Buildpack

## `ghcr.io/h2ik/buildpack-sidekiq`

The Sidekiq CNB sets the start command for a given ruby application that runs on a [sidekiq server](https://sidekiq.org/).

## Integration

This CNB writes a start command, so there's currently no scenario we can
imagine that you would need to require it as dependency. If a user likes to
include some other functionality, it can be done independent of the Sidekiq CNB
without requiring a dependency of it.

To package this buildpack for consumption:
```
$ ./scripts/package.sh
```
This builds the buildpack's source using GOOS=linux by default. You can supply another value as the first argument to package.sh.

## `buildpack.yml` Configurations

There are no extra configurations for this buildpack based on `buildpack.yml`.

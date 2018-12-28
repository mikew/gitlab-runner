# gitlab-runner mods

Two simple changes so we don't have to muck with `cache`:

- `git checkout -- .` instead of `git clean -ffdx && git reset --hard`
- Builds happen in a directory like `ci--{namespace}--{project}--{branch}`

## Building

- Install Go
- Install xz
- Install Docker
- Cherry-pick the commits from a previous `*-mods` branch
- Run `BUILD_PLATFORMS="-osarch 'darwin/amd64 linux/amd64 windows/amd64'" make deps build`

## Installing

- Open `out/binaries/`
  - Linux: Move to `/usr/bin/gitlab-runner` and `/usr/bin/gitlab-ci-multi-runner` (see sst-rickilake/sst-ci-runner/docker-compose.yml)
  - macOS: Move to `/usr/local/bin/gitlab-ci-multi-runner`
  - Windows: Move to `%USERPROFILE%\Desktop\gitlab-ci-runner\gitlab-ci-multi-runner.exe`
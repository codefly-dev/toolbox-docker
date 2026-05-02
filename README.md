# toolbox-docker

A codefly toolbox plugin for Docker image and container inspection.
Canonical owner of the `docker` binary — the `codefly-dev/toolbox-bash`
plugin refuses every `docker` invocation and routes callers here.

## Tools (read-only)

- `docker.list_containers(all?)` — running containers (or all w/ stopped).
  Returns `{containers: [{id, image, status, state, names}, ...]}`.
- `docker.list_images()` — every image cached in the local daemon:
  id, repo_tags, size, created_unix.
- `docker.inspect_container(id)` — curated subset of `docker inspect`:
  id, name, image, created, running, status, exit_code.

Mutating operations (`run`, `stop`, `rm`, `exec`) are deliberately
omitted from Phase 1. They land in a later iteration once the
sandbox + policy story for "agent runs containers" is settled.

## Sandbox

The Docker daemon is reached via `/var/run/docker.sock`. The plugin's
declared sandbox grants access to that one Unix socket, denies network
+ filesystem reads/writes by default. Plugins running under bwrap
(Linux) or sandbox-exec (macOS) inherit those constraints.

## Configuration

| Env var                    | Default       | Purpose                                       |
| -------------------------- | ------------- | --------------------------------------------- |
| `CODEFLY_TOOLBOX_VERSION`  | `0.0.0-dev`   | Identity version surfaced via `Identity()`    |

The Docker SDK reads `DOCKER_HOST` / `DOCKER_TLS_VERIFY` /
`DOCKER_CERT_PATH` from the plugin's environment as usual.

## Build & test

```bash
go build ./...
go test ./...
```

## Contract

This plugin implements the codefly Toolbox gRPC contract defined in
[`codefly-dev/core`](https://github.com/codefly-dev/core) at
`proto/codefly/services/toolbox/v0/toolbox.proto`.

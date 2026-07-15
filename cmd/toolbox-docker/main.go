// Command toolbox-docker is the standalone binary form of the codefly
// docker toolbox plugin. Loaded via the standard agent loader
// (core/agents/manager.Load); registers a Toolbox server through
// agents.Serve.
//
// Configuration:
//
//	CODEFLY_TOOLBOX_VERSION — Identity version. Default "0.0.0-dev".
//	(The Docker daemon is discovered from the standard DOCKER_HOST/
//	DOCKER_TLS_VERIFY/DOCKER_CERT_PATH env vars via the Docker SDK.)
package main

import (
	"github.com/codefly-dev/core/agents"
	coretoolbox "github.com/codefly-dev/core/toolbox"
	docker "github.com/codefly-dev/toolbox-docker"
)

func main() {
	agents.ServeToolbox(docker.New(coretoolbox.Version()))
}

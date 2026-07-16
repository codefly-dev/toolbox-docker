package docker_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/codefly-dev/core/resources"
	docker "github.com/codefly-dev/toolbox-docker"
)

func TestManifestMatchesProductionCatalog(t *testing.T) {
	manifest, err := resources.LoadToolboxFromDir(context.Background(), ".")
	require.NoError(t, err)
	require.NoError(t, manifest.ValidateForProduction())

	server := docker.New(manifest.Version)
	defer server.Close()
	names := make([]string, 0, len(server.Tools()))
	for _, tool := range server.Tools() {
		names = append(names, tool.Name)
	}
	require.NoError(t, manifest.ValidateToolCatalog(names...))
}

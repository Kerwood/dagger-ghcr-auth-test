// A generated module for AuthTest functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/auth-test/internal/dagger"
)

type AuthTest struct{}

// Returns lines that match a pattern in the files of the provided Directory
func (m *AuthTest) PublishImage(ctx context.Context, ghUser string, ghToken *dagger.Secret) (string, error) {

	return dag.Container().
		From("alpine:latest").
		WithLabel("org.opencontainers.image.source", "https://github.com/kerwood/dagger-ghcr-auth-test").
		WithRegistryAuth("ghcr.io", ghUser, ghToken).
		Publish(ctx, "ghcr.io/kerwood/dagger-ghcr-auth-test:latest")
}

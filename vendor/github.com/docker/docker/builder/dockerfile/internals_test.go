package dockerfile

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/builder"
	"github.com/docker/docker/builder/remotecontext"
	"github.com/docker/docker/pkg/archive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmptyDockerfile(t *testing.T) {
	contextDir, cleanup := createTestTempDir(t, "", "builder-dockerfile-test")
	defer cleanup()

	createTestTempFile(t, contextDir, builder.DefaultDockerfileName, "", 0777)

	readAndCheckDockerfile(t, "emptyDockerfile", contextDir, "", "the Dockerfile (Dockerfile) cannot be empty")
}

func TestSymlinkDockerfile(t *testing.T) {
	contextDir, cleanup := createTestTempDir(t, "", "builder-dockerfile-test")
	defer cleanup()

	createTestSymlink(t, contextDir, builder.DefaultDockerfileName, "/etc/passwd")

	// The reason the error is "Cannot locate specified Dockerfile" is because
	// in the builder, the symlink is resolved within the context, therefore
	// Dockerfile -> /etc/passwd becomes etc/passwd from the context which is
	// a nonexistent file.
	expectedError := fmt.Sprintf("Cannot locate specified Dockerfile: %s", builder.DefaultDockerfileName)

	readAndCheckDockerfile(t, "symlinkDockerfile", contextDir, builder.DefaultDockerfileName, expectedError)
}

func TestDockerfileOutsideTheBuildContext(t *testing.T) {
	contextDir, cleanup := createTestTempDir(t, "", "builder-dockerfile-test")
	defer cleanup()

	expectedError := "Forbidden path outside the build context: ../../Dockerfile ()"

	readAndCheckDockerfile(t, "DockerfileOutsideTheBuildContext", contextDir, "../../Dockerfile", expectedError)
}

func TestNonExistingDockerfile(t *testing.T) {
	contextDir, cleanup := createTestTempDir(t, "", "builder-dockerfile-test")
	defer cleanup()

	expectedError := "Cannot locate specified Dockerfile: Dockerfile"

	readAndCheckDockerfile(t, "NonExistingDockerfile", contextDir, "Dockerfile", expectedError)
}

func readAndCheckDockerfile(t *testing.T, testName, contextDir, dockerfilePath, expectedError string) {
	tarStream, err := archive.Tar(contextDir, archive.Uncompressed)
	require.NoError(t, err)

	defer func() {
		if err = tarStream.Close(); err != nil {
			t.Fatalf("Error when closing tar stream: %s", err)
		}
	}()

	if dockerfilePath == "" { // handled in BuildWithContext
		dockerfilePath = builder.DefaultDockerfileName
	}

	_, _, err = remotecontext.Detect(context.Background(), "", dockerfilePath, tarStream, nil)
	assert.EqualError(t, err, expectedError)
}

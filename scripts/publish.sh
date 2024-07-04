#!/bin/sh

set -e

# Get the last version in repository
CURRENT_VERSION=$(git describe --abbrev=0 --tags)

# Get the declared version in .version file
VERSION=$(cat .version)

# Check if versions are equal
if [ "$CURRENT_VERSION" = "$VERSION" ]; then
    echo "Versions are equal. Nothing to publish"
    exit 0
fi

# Tidy the repository
go mod tidy

# Create the new tag version
git tag "$VERSION"

# Push the new tag version
git push origin "$VERSION"

# Wait for the new version to be published
echo "Waiting for pkg.go.dev to publish the new version..."
sleep 10

# Publish the new version
GOPROXY=proxy.golang.org go list -m github.com/Diegiwg/gogui@"$VERSION" &

# Ensure that the new version was published
echo "Go to https://pkg.go.dev/github.com/Diegiwg/gogui@$VERSION to check the new version"

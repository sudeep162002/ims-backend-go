#!/bin/bash

VERSION=""

# Parse command-line options
while getopts v: flag; do
  case "${flag}" in
    v) VERSION=${OPTARG};;
  esac
done

# Fetch tags and prune
git fetch --tags --prune --unshallow 2>/dev/null

# Get the highest tag number and set initial version if not found
CURRENT_VERSION=$(git describe --abbrev=0 --tags 2>/dev/null || echo "v1.1.1")
echo "Current Version: $CURRENT_VERSION"

# Parse version components
IFS='.' read -r -a CURRENT_VERSION_PARTS <<<"${CURRENT_VERSION//v}"
VNUM1=${CURRENT_VERSION_PARTS[0]}
VNUM2=${CURRENT_VERSION_PARTS[1]}
VNUM3=${CURRENT_VERSION_PARTS[2]}

# Increment version based on parameter
case "$VERSION" in
  major) ((VNUM1++)); VNUM2=0; VNUM3=0;;
  minor) ((VNUM2++)); VNUM3=0;;
  patch) ((VNUM3++));;
  *)
    echo "Invalid version type. Use: -v [major, minor, patch]"
    exit 1
    ;;
esac

# Create new tag
NEW_TAG="v$VNUM1.$VNUM2.$VNUM3"
echo "($VERSION) Updating $CURRENT_VERSION to $NEW_TAG"

# Check if the current commit is already tagged
GIT_COMMIT=$(git rev-parse HEAD)
NEEDS_TAG=$(git describe --contains "$GIT_COMMIT" 2>/dev/null)

# Tag the current commit if no tag exists
if [ -z "$NEEDS_TAG" ]; then
  echo "Tagged with $NEW_TAG"
  git tag "$NEW_TAG"
  git push --tags
  git push
else
  echo "Already a tag on this commit"
fi

# Set GitHub Actions output
echo ::set-output name=git-tag::$NEW_TAG

exit 0

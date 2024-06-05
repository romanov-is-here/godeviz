#!/bin/bash

set -euo pipefail

if [ -z "${RELEASE_VERSION+x}" ]; then
  echo "Release version \$RELEASE_VERSION is not set."
  exit 1
fi

git checkout -b release/"v${RELEASE_VERSION}"
git push origin release/"v${RELEASE_VERSION}"
git tag -a "v$RELEASE_VERSION" -m "Release v$RELEASE_VERSION"
git push origin "v$RELEASE_VERSION"
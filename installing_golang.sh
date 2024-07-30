#!/usr/bin/env bash

# let us download a file with curl on Linux command line #
VERSION="1.22.5" # go version
ARCH="amd64"     # go archicture
curl -O -L "https://golang.org/dl/go${VERSION}.linux-${ARCH}.tar.gz"

# Instead of curl, one can use wget command too #
# wget -L "https://golang.org/dl/go${VERSION}.linux-${ARCH}.tar.gz"

ls -l

#!/bin/bash

# Build client binary
#
go build -o binaries/co-client src/cmd/client/main.go src/cmd/client/structures.go
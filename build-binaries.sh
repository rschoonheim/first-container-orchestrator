#!/bin/bash

# Build the control plane
#
go build -o binaries/co-control-plane src/cmd/control_plane/main.go

# Build the storage server binary
#
go build -o binaries/co-storage src/cmd/storage/main.go

# Build the API server binary
#
go build -o binaries/co-api src/cmd/api/main.go

# Build the CLI client binary
#
go build -o binaries/co-client src/cmd/client/main.go src/cmd/client/structures.go
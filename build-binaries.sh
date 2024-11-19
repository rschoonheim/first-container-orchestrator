#!/bin/bash

# Build the storage server binary
#
go build -o binaries/co-storage src/cmd/storage/main.go

# Build the API server binary
#
go build -o binaries/co-api src/cmd/api/main.go

# Build the CLI client binary
#
go build -o binaries/co-client src/cmd/client/main.go src/cmd/client/structures.go
package main

import (
	"context"
	network_v1 "first-container-orchestrator/network/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log/slog"
	"os"
)

func main() {

	// gRPC Client
	//
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		slog.Error("Failed to connect to server", "error", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Create network client
	//
	networkClient := network_v1.NewNetworkServiceClient(conn)

	uuidString := uuid.New().String()
	ctx := context.Background()

	// Create network
	//
	network, err := networkClient.CreateNetwork(ctx, &network_v1.CreateNetworkRequest{Id: uuidString})
	if err != nil {
		slog.Error("Failed to create network", "error", err)
		os.Exit(1)
	}

	slog.Info("Network", "network", network)

	// Get network
	//
	network, err = networkClient.GetNetwork(ctx, &network_v1.GetNetworkRequest{Id: uuidString})
	if err != nil {
		slog.Error("Failed to get network", "error", err)
	}

	slog.Info("Network", "network", network)
}

package network_v1

import (
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"os"
)

// storageGet - Get a storage by its ID.
func storageGet(id string) (*Network, error) {

	// Check file system of the exists of "networks/{id}" directory
	//
	if _, err := os.Stat("storage/networks/" + id); os.IsNotExist(err) {
		return nil, nil
	}

	// Get state from file
	//
	path := "storage/networks/" + id + "/state.json"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, status.Errorf(codes.Internal, "INTERNAL_ERROR")
	}

	// Read state from file
	//
	file, err := os.Open(path)
	if err != nil {
		slog.Error("Failed to read state file of network", "error", err)
		return nil, status.Errorf(codes.Internal, "INTERNAL_ERROR")
	}
	defer file.Close()

	// Decode state
	//
	network := &Network{}
	if err := json.NewDecoder(file).Decode(network); err != nil {
		slog.Error("Failed to decode state file of network", "error", err)
		return nil, status.Errorf(codes.Internal, "INTERNAL_ERROR")
	}

	return network, nil
}

// storageCreate - Creates a storage.
func storageCreate(network *Network) (*Network, error) {
	// Check if the network ID is unique
	//
	if _, err := os.Stat("storage/networks/" + network.GetId()); !os.IsNotExist(err) {
		return nil, status.Errorf(codes.AlreadyExists, "NETWORK_ALREADY_EXISTS")
	}

	// Create the network directory
	//
	if err := os.Mkdir("storage/networks/"+network.GetId(), 0755); err != nil {
		return nil, status.Errorf(codes.Internal, "INTERNAL_ERROR")
	}

	// Write network to state file
	//
	path := "storage/networks/" + network.GetId() + "/state.json"
	file, err := os.Create(path)
	if err != nil {
		slog.Error("Failed to create state file of network", "error", err)
		return nil, status.Errorf(codes.Internal, "INTERNAL_ERROR")
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(network); err != nil {
		slog.Error("Failed to encode state file of network", "error", err)
		return nil, status.Errorf(codes.Internal, "INTERNAL_ERROR")
	}

	return network, nil
}

// storageNetworkNameAvailable - Check if a network name is available.
func storageNetworkNameAvailable(name string) bool {
	// Check file system of the exists of "networks/{name}" directory
	//
	if _, err := os.Stat("storage/networks/" + name); os.IsNotExist(err) {
		return true
	}

	return false
}

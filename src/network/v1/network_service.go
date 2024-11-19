package network_v1

import (
	"context"
	"first-container-orchestrator/internal/validation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NetworkService struct{}

// GetNetwork - Get a network by its ID.
func (n *NetworkService) GetNetwork(ctx context.Context, in *GetNetworkRequest) (*Network, error) {

	// Make sure that the incoming request is valid.
	// ------------------------------------------------------------------------
	//
	//
	if validation.IsUuid(in.GetId()) == false {
		return nil, status.Errorf(codes.InvalidArgument, "ID_MALFORMED")
	}

	// Get the network from storage
	// ------------------------------------------------------------------------
	//
	//
	network, err := storageGet(in.GetId())
	if err != nil {
		return nil, err
	}

	if network == nil {
		return nil, status.Errorf(codes.NotFound, "NETWORK_NOT_FOUND")
	}

	return network, nil
}

// CreateNetwork - Creates a network.
func (n *NetworkService) CreateNetwork(ctx context.Context, in *CreateNetworkRequest) (*Network, error) {

	// Make sure that the incoming request is valid.
	// ------------------------------------------------------------------------
	//
	//

	if validation.IsUuid(in.GetId()) == false {
		return nil, status.Errorf(codes.InvalidArgument, "ID_MALFORMED")
	}

	if in.GetName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "NAME_NOT_PROVIDED")
	}

	if storageNetworkNameAvailable(in.GetName()) == false {
		return nil, status.Errorf(codes.AlreadyExists, "NAME_NOT_AVAILABLE")
	}

	// Is the network ID unique?
	uniqueId, _ := storageGet(in.GetId())
	if uniqueId != nil {
		return nil, status.Errorf(codes.AlreadyExists, "ID_NOT_UNIQUE")
	}

	// Initialize the network creation process.
	// ------------------------------------------------------------------------
	//
	//

	network, err := storageCreate(&Network{
		Id: in.GetId(),
	})
	if err != nil {
		return nil, err
	}

	network, err = networkImplement(network)
	if err != nil {
		return nil, err
	}

	return network, nil
}

func (n *NetworkService) mustEmbedUnimplementedNetworkServiceServer() {}

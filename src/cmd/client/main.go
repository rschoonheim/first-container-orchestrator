package main

import (
	"first-container-orchestrator/internal/cni"
	"github.com/google/uuid"
	"log/slog"
	"os"
)

func main() {

	// Delete the network namespace before exiting..
	//
	cni.NetworkNamespaceDelete(&cni.NetworkNamespace{
		Name: "test",
	})

	cni.NetworkNamespaceDelete(&cni.NetworkNamespace{
		Name: "test2",
	})
	networkNamespace, err := cni.NetworkNamespaceCreate(
		&cni.NetworkNamespace{
			Name: uuid.New().String(),
		},
	)

	if err != nil {
		slog.Error("Failed to create network namespace", "error", err)
		os.Exit(1)
	}

	networkNamespaceTwo, err := cni.NetworkNamespaceCreate(
		&cni.NetworkNamespace{
			Name: uuid.New().String(),
		},
	)

	if err != nil {
		slog.Error("Failed to create network namespace", "error", err)
		os.Exit(1)
	}

	// Add virtual cable between network namespaces
	//
	vcable, err := cni.NetworkVirtualCableCreate(&cni.NetworkVirtualCable{
		SourceNetworkNamespace:      networkNamespace,
		DestinationNetworkNamespace: networkNamespaceTwo,
	})

	if err != nil {
		slog.Error("Failed to create network virtual cable", "error", err)
		os.Exit(1)
	}

	slog.Info("Network virtual cable created", "cable", vcable)

	// Delete the network namespace before exiting..
	//
	cni.NetworkNamespaceDelete(&cni.NetworkNamespace{
		Name: "test",
	})

	cni.NetworkNamespaceDelete(&cni.NetworkNamespace{
		Name: "test2",
	})

	os.Exit(0)

}

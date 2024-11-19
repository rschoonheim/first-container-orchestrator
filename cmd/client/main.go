package main

import (
	"first-container-orchestrator/internal/cni"
	"log/slog"
	"os"
)

func main() {

	networkNamespace, err := cni.NetworkNamespaceCreate(
		&cni.NetworkNamespace{
			Name: "test",
		},
	)

	if err != nil {
		slog.Error("Failed to create network namespace", "error", err)
		os.Exit(1)
	}

	// Delete the network namespace before exiting..
	//
	cni.NetworkNamespaceDelete(&cni.NetworkNamespace{
		Name: "test",
	})

	println("Network namespace created", networkNamespace.Name)
	os.Exit(0)

}

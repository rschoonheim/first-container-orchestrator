package cni

import (
	"errors"
	"log/slog"
	"os/exec"
	"regexp"
)

type NetworkNamespace struct {
	// Name - The name of the network namespace.
	Name string `json:"name"`
}

// NetworkNamespaceCreate - Creates a network namespace on the host machine.
func NetworkNamespaceCreate(networkNamespace *NetworkNamespace) (*NetworkNamespace, error) {

	// Steps to create a network namespace
	// ------------------------------------------------------------------------
	// 1. Check if the network namespace name is unique.
	// 2. Create the network namespace.
	//

	// Check if the network namespace name is unique
	//
	cmd := exec.Command("ip", "netns", "list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	nameExistsRegexPattern := regexp.MustCompile(`(?m)^` + networkNamespace.Name + `$`)
	if nameExistsRegexPattern.MatchString(string(output)) {
		slog.Debug("Tried to create network namespace with a name that already exists", "name", networkNamespace.Name)
		return nil, errors.New("NETWORK_NAMESPACE_NAME_ALREADY_EXISTS")
	}

	// Create the network namespace
	//
	cmd = exec.Command("ip", "netns", "add", networkNamespace.Name)
	output, err = cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New("FAILED_TO_CREATE_NETWORK_NAMESPACE")
	}

	return networkNamespace, nil
}

// NetworkNamespaceDelete - Deletes a network namespace on the host machine.
func NetworkNamespaceDelete(networkNamespace *NetworkNamespace) error {

	// Steps to delete a network namespace
	// ------------------------------------------------------------------------
	// 1. Check if the network namespace exists.
	// 2. Delete the network namespace.
	//

	// Check if the network namespace exists
	//
	cmd := exec.Command("ip", "netns", "list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	nameExistsRegexPattern := regexp.MustCompile(`(?m)^` + networkNamespace.Name + `$`)
	if !nameExistsRegexPattern.MatchString(string(output)) {
		slog.Debug("Tried to delete network namespace that does not exist", "name", networkNamespace.Name)
		return errors.New("NETWORK_NAMESPACE_DOES_NOT_EXIST")
	}

	// Delete the network namespace
	//
	cmd = exec.Command("ip", "netns", "delete", networkNamespace.Name)
	output, err = cmd.CombinedOutput()
	if err != nil {
		return errors.New("FAILED_TO_DELETE_NETWORK_NAMESPACE")
	}

	return nil
}

package cni

import (
	"errors"
	"github.com/google/uuid"
	"os"
	"os/exec"
	"strings"
)

type NetworkVirtualCableInterface struct {
	// Id - The name of the network virtual cable interface.
	Id string `json:"id"`

	// IpAddress - The IP address of the network virtual cable interface.
	IpAddress string `json:"ip_address"`
}

type NetworkVirtualCable struct {
	// Id - The ID of the network virtual cable.
	Id string `json:"id"`

	// Veth0 - The name of the first veth pair.
	Veth0 *NetworkVirtualCableInterface `json:"veth0"`

	// Veth1 - The name of the second veth pair.
	Veth1 *NetworkVirtualCableInterface `json:"veth1"`

	// SourceNetworkNamespace - The source network namespace.
	SourceNetworkNamespace *NetworkNamespace `json:"source_network_namespace"`

	// DestinationNetworkNamespace - The destination network namespace.
	DestinationNetworkNamespace *NetworkNamespace `json:"destination_network_namespace"`
}

var (
	existingNetworkVirtualCables = make(map[string]*NetworkVirtualCable)
)

func init() {
	os.MkdirAll("storage/cni/network_virtual_cables", 0755)
}

// NetworkVirtualCableExists - Checks if a network virtual cable exists.
func NetworkVirtualCableExists(networkVirtualCable *NetworkVirtualCable) bool {
	for _, existingNetworkVirtualCable := range existingNetworkVirtualCables {
		if existingNetworkVirtualCable.SourceNetworkNamespace.Name == networkVirtualCable.SourceNetworkNamespace.Name &&
			existingNetworkVirtualCable.DestinationNetworkNamespace.Name == networkVirtualCable.DestinationNetworkNamespace.Name {
			return true
		}
	}

	return false
}

// networkVirtualCableGenerateUniqueVethNames - Generates unique names for the veth pair.
func networkVirtualCableGenerateUniqueVethNames() (string, string) {
	var randId = uuid.New().String()[:11]
	randId = strings.ReplaceAll(randId, "-", "")
	return "veth" + randId + "0", "veth" + randId + "1"
}

// NetworkVirtualCableCreate - Creates a network virtual cable between two network namespaces.
func NetworkVirtualCableCreate(networkVirtualCable *NetworkVirtualCable) (*NetworkVirtualCable, error) {
	var veth0, veth1 = networkVirtualCableGenerateUniqueVethNames()
	var cable = &NetworkVirtualCable{
		Id:                          uuid.New().String(),
		Veth0:                       &NetworkVirtualCableInterface{Id: veth0, IpAddress: "192.168.64.1/24"},
		Veth1:                       &NetworkVirtualCableInterface{Id: veth1, IpAddress: "192.168.64.2/24"},
		SourceNetworkNamespace:      networkVirtualCable.SourceNetworkNamespace,
		DestinationNetworkNamespace: networkVirtualCable.DestinationNetworkNamespace,
	}

	if NetworkVirtualCableExists(cable) {
		return nil, errors.New("NETWORK_VIRTUAL_CABLE_ALREADY_EXISTS")
	}

	// Initialize the storage for the virtual cable
	//
	storageVirtualNetworkCableInit(cable)

	// Generate unique names for the veth pair
	//
	_, err := exec.Command("ip", "link", "add", cable.Veth0.Id, "type", "veth", "peer", "name", cable.Veth1.Id).Output()
	if err != nil {
		return nil, errors.New("FAILED_TO_CREATE_VETH_PAIR")
	}

	// Set the veth pair to the network namespaces
	//
	_, err = exec.Command("ip", "link", "set", cable.Veth0.Id, "netns", cable.SourceNetworkNamespace.Name).Output()
	if err != nil {
		return nil, errors.New("FAILED_TO_SET_VETH0_TO_SOURCE_NAMESPACE")
	}

	_, err = exec.Command("ip", "link", "set", cable.Veth1.Id, "netns", cable.DestinationNetworkNamespace.Name).Output()
	if err != nil {
		return nil, errors.New("FAILED_TO_SET_VETH1_TO_DESTINATION_NAMESPACE")
	}

	// Add ip addresses to the veth pair
	//
	_, err = exec.Command("ip", "netns", "exec", cable.SourceNetworkNamespace.Name, "ip", "addr", "add", cable.Veth0.IpAddress, "dev", cable.Veth0.Id).Output()
	if err != nil {
		return nil, errors.New("FAILED_TO_ADD_IP_ADDRESS_TO_VETH0")
	}

	_, err = exec.Command("ip", "netns", "exec", cable.DestinationNetworkNamespace.Name, "ip", "addr", "add", cable.Veth1.IpAddress, "dev", cable.Veth1.Id).Output()
	if err != nil {
		return nil, errors.New("FAILED_TO_ADD_IP_ADDRESS_TO_VETH1")
	}

	// Bring the veth pair up in their namespace
	//
	_, err = exec.Command("ip", "netns", "exec", cable.SourceNetworkNamespace.Name, "ip", "link", "set", cable.Veth0.Id, "up").Output()
	if err != nil {
		return nil, errors.New("FAILED_TO_BRING_VETH0_UP")
	}

	_, err = exec.Command("ip", "netns", "exec", cable.DestinationNetworkNamespace.Name, "ip", "link", "set", cable.Veth1.Id, "up").Output()
	if err != nil {
		return nil, errors.New("FAILED_TO_BRING_VETH1_UP")
	}

	return cable, nil
}

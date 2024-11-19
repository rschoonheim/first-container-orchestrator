package cni

import (
	"encoding/json"
	"os"
)

// storageVirtualNetworkCableGetPath - Gets the path to the storage for a virtual network cable.
func storageVirtualNetworkCableGetPath(id string) string {
	return "storage/cni/network_virtual_cables/" + id
}

// storageVirtualNetworkCableInit - Initializes the storage for a virtual network cable.
func storageVirtualNetworkCableInit(cable *NetworkVirtualCable) {
	os.MkdirAll(storageVirtualNetworkCableGetPath(cable.Id), 0755)

	// Save the virtual network cable to storage
	//
	storageVirtualNetworkCableSave(cable)
}

// storageVirtualNetworkCableSave - Saves a virtual network cable to storage.
func storageVirtualNetworkCableSave(cable *NetworkVirtualCable) {
	// Save the virtual network cable to storage
	//
	file, _ := os.Create(storageVirtualNetworkCableGetPath(cable.Id) + "/state.json")
	defer file.Close()

	_ = json.NewEncoder(file).Encode(cable)

}

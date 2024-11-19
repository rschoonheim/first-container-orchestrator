# Container Network Interface (CNI)

The CNI abstracts the network configuration for a container runtime.

## Files

### network_namespace

This file contains functions to create a network namespace

### network_virtual_cable

This file contains functions to create a virtual cable between two network namespaces. A virtual cable is a pair of
veth devices that are connected to each other. One end of the cable is in the source network namespace and the other
end is in the destination network namespace.
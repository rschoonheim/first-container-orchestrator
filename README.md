# First Container Orchestrator

My first attempt at building a container orchestrator. This project is a learning experience and is not intended to be
used in production.

## Systems

![System overview](docs/images/system_overview.png)

The system is composed of four main components: the control plane, the storage server, the API server, and the client.
Each component has a specific role in the system.

### Control Plane
The control plane is responsible for synchronizing the state of Linux with the state of the system. 

### Storage Server
The storage server is responsible for managing the state of the system. 

### API Server
The API server is a RESTful API providing an interface to the system.

### Command Line Interface (CLI)
The CLI is a command line interface that interacts with the API server.

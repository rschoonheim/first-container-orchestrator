# Mutual TLS

This module provides a simple interface for creating a mutual TLS connection between a client and a server.

### What is Mutual TLS?

Mutual TLS, or mTLS for short, is a method for mutual authentication. mTLS ensures that the parties at each end of a
network connection are who they claim to be by verifying that they both have the correct private key. The information
within their respective TLS certificates provides additional verification.

[More information about mTLS](https://www.cloudflare.com/learning/access-management/what-is-mutual-tls)

### Why is Mutual TLS needed in this project?

There will be communication between different components of the system. To ensure that only authorized clients can 
communicate on the network, mTLS is used to ensure that the client is who they claim to be.


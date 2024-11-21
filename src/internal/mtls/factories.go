package mtls

// MakeResource - Create a new mTLS resource.
func MakeResource(cert []byte, key []byte, ca []byte) *MTlsResource {
	return &MTlsResource{
		Cert:                 cert,
		Key:                  key,
		CertificateAuthority: ca,
	}
}

// MakeConnection - Create a new mTLS connection.
func MakeConnection(client *MTlsResource, server *MTlsResource) *MTlsConnection {
	return &MTlsConnection{
		Client: client,
		Server: server,
	}
}

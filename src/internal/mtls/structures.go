package mtls

// MTlsResource - A mTLS resource can be a client or server.
type MTlsResource struct {
	// Cert - The certificate of the server.
	Cert []byte

	// Key - The private key of the server.
	Key []byte

	// CertificateAuthority - The certificate authority to validate the client certificate.
	CertificateAuthority []byte
}

type MTlsConnection struct {
	// Client - The client side of the mTLS connection.
	Client *MTlsResource

	// Server - The server side of the mTLS connection.
	Server *MTlsResource
}

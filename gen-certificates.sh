#!/bin/bash

rm -rf certs

# Make certs directory
mkdir -p certs

# Generate (self-signed) root certificate authority
openssl genrsa -out certs/ca.key 2048
openssl req -new -x509 -days 365 -key certs/ca.key -out certs/ca.crt -subj "/C=NL/ST=Utrecht/L=Utrecht/O=FCO Authority/CN=FCO Authority"

# Generate intermediate certificate authority for the server
openssl genrsa -out certs/server-ca.key 2048
openssl req -new -key certs/server-ca.key -out certs/server-ca.csr -subj "/C=NL/ST=Utrecht/L=Utrecht/O=FCO Server CA/CN=FCO Server CA"
openssl x509 -req -days 365 -in certs/server-ca.csr -CA certs/ca.crt -CAkey certs/ca.key -CAcreateserial -out certs/server-ca.crt -extfile <(printf "basicConstraints=critical,CA:TRUE\nkeyUsage=critical,keyCertSign,cRLSign")

openssl verify -CAfile certs/ca.crt certs/server-ca.crt

# Create the full chain
cat certs/server-ca.crt certs/ca.crt > certs/server-ca-chain.crt
openssl verify -CAfile certs/ca.crt certs/server-ca-chain.crt

# Generate server certificate signed by the intermediate certificate authority
openssl genrsa -out certs/server.key 2048
openssl req -new -key certs/server.key -out certs/server.csr -subj "/C=NL/ST=Utrecht/L=Utrecht/O=FCO Server/CN=localhost"
openssl x509 -req -days 365 -in certs/server.csr -CA certs/server-ca.crt -CAkey certs/server-ca.key -CAcreateserial -out certs/server.crt -extfile <(printf "subjectAltName=DNS:localhost,IP:127.0.0.1")

# Verify the server certificate using the full chain
openssl verify -CAfile certs/server-ca-chain.crt certs/server.crt

# Generate client certificate signed by the intermediate certificate authority
openssl genrsa -out certs/client.key 2048
openssl req -new -key certs/client.key -out certs/client.csr -subj "/C=NL/ST=Utrecht/L=Utrecht/O=FCO Client/CN=client"
openssl x509 -req -days 365 -in certs/client.csr -CA certs/server-ca.crt -CAkey certs/server-ca.key -CAcreateserial -out certs/client.crt

# Verify the client certificate using the full chain
openssl verify -CAfile certs/server-ca-chain.crt certs/client.crt

# Remove all certificate signing requests
rm certs/*.csr
mkdir -p certs

# Install mkcert trusty certificate in the OS
mkcert -ecdsa -install

# Generate public - private key 
go run filippo.io/mkcert -ecdsa -cert-file "certs/certificate.pem" -key-file "certs/certificate.key" localhost 127.0.0.1 ::1
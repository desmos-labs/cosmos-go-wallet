package types

import (
	"crypto/tls"
	"regexp"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	HTTPProtocols = regexp.MustCompile("https?://")
)

// CreateGrpcConnection creates a new gRPC client connection from the given configuration
func CreateGrpcConnection(address string) (*grpc.ClientConn, error) {
	var transportCredentials credentials.TransportCredentials
	if strings.HasPrefix(address, "https") {
		transportCredentials = credentials.NewTLS(&tls.Config{MinVersion: tls.VersionTLS12})
	} else {
		transportCredentials = insecure.NewCredentials()
	}

	address = HTTPProtocols.ReplaceAllString(address, "")
	return grpc.Dial(address, grpc.WithTransportCredentials(transportCredentials))
}

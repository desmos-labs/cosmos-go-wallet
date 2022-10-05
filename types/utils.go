package types

import (
	"crypto/tls"
	"regexp"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	HTTPProtocols = regexp.MustCompile("https?://")
)

// CreateGrpcConnection creates a new gRPC client connection from the given configuration
func CreateGrpcConnection(address string) (*grpc.ClientConn, error) {
	var grpcOpts []grpc.DialOption
	if strings.HasPrefix(address, "https") {
		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			MinVersion: tls.VersionTLS12,
		})))
	} else {
		grpcOpts = append(grpcOpts, grpc.WithInsecure())
	}

	address = HTTPProtocols.ReplaceAllString(address, "")
	return grpc.Dial(address, grpcOpts...)
}

package main

import (
	"log"
	"net"

	pd "github.com/wangyuche/time-manager/proto/timemanager"
	"google.golang.org/grpc"
)

func init() {

}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pd.RegisterRouteGuideServer(grpcServer, nil)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

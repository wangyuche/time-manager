package main

import (
	"context"
	"log"
	"net"

	tm "github.com/wangyuche/time-manager/proto/timemanager"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func init() {

}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	tm.RegisterTimeGRPCServer(grpcServer, &TimeGRPC{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

type TimeGRPC struct {
	tm.UnimplementedTimeGRPCServer
	offset  int
	accrate int
}

func (s *TimeGRPC) GetServerTime(ctx context.Context, void *emptypb.Empty) (*tm.GetServerTimeRes, error) {
	return &tm.GetServerTimeRes{Offset: int32(s.offset), Accrate: int32(s.accrate)}, nil
}

func (s *TimeGRPC) SetServerTime(ctx context.Context, req *tm.SetServerTimeReq) (*tm.SetServerTimeRes, error) {
	s.offset = int(req.Offset)
	s.accrate = int(req.Accrate)
	return &tm.SetServerTimeRes{Status: tm.TimeStatus_success}, nil
}

func (s *TimeGRPC) ClearServerTime(ctx context.Context, void *emptypb.Empty) (*tm.ClearServerTimeRes, error) {
	return &tm.ClearServerTimeRes{Status: tm.TimeStatus_success}, nil
}

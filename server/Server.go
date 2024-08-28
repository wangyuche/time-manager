package main

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

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
	tm.RegisterTimeGRPCServer(grpcServer, &TimeGRPC{broadcast: make(map[chan bool]bool), m: new(sync.RWMutex)})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

type TimeGRPC struct {
	tm.UnimplementedTimeGRPCServer
	offset    int
	accrate   int
	timeunix  int64
	broadcast map[chan bool]bool
	m         *sync.RWMutex
}

func (s *TimeGRPC) GetServerTime(ctx context.Context, void *emptypb.Empty) (*tm.GetServerTimeRes, error) {
	return &tm.GetServerTimeRes{Timeunix: int64(s.timeunix), Offset: int32(s.offset), Accrate: int32(s.accrate)}, nil
}

func (s *TimeGRPC) SetServerTime(ctx context.Context, req *tm.SetServerTimeReq) (*tm.SetServerTimeRes, error) {
	s.offset = int(req.Offset)
	s.accrate = int(req.Accrate)
	if s.timeunix == 0 {
		s.timeunix = time.Now().Unix()
	}
	for b, _ := range s.broadcast {
		s.m.RLock()
		b <- true
		s.m.RUnlock()
	}
	return &tm.SetServerTimeRes{Status: tm.TimeStatus_success}, nil
}

func (s *TimeGRPC) ClearServerTime(ctx context.Context, void *emptypb.Empty) (*tm.ClearServerTimeRes, error) {
	s.accrate = 0
	s.offset = 0
	s.timeunix = 0
	for b, _ := range s.broadcast {
		s.m.RLock()
		b <- true
		s.m.RUnlock()
	}
	return &tm.ClearServerTimeRes{Status: tm.TimeStatus_success}, nil
}

func (s *TimeGRPC) BroadcastServerTime(void *emptypb.Empty, stream tm.TimeGRPC_BroadcastServerTimeServer) error {
	ch := make(chan bool)
	s.broadcast[ch] = false
	for {
		<-ch
		err := stream.Send(&tm.GetServerTimeRes{Timeunix: int64(s.timeunix), Offset: int32(s.offset), Accrate: int32(s.accrate)})
		if err != nil {
			log.Printf("Error: %v", err)
			s.m.Lock()
			delete(s.broadcast, ch)
			s.m.Unlock()
			return err
		}
	}
}

package TimeManager

import (
	"context"
	"io"
	"log"
	"time"

	tm "github.com/wangyuche/time-manager/proto/timemanager"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var address string = "localhost:9000"
var conn *grpc.ClientConn
var offset int32
var accrate int32
var timer tm.TimeGRPCClient
var cachetime int64

func init() {
	go func() {
		for {
			if err := connect(); err != nil {
				log.Printf("Error: %v", err)
				time.Sleep(5 * time.Second)
			}
		}
	}()
}

func connect() error {
	var err error
	conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	timer = tm.NewTimeGRPCClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	r, err := timer.GetServerTime(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}
	accrate = r.Accrate
	offset = r.Offset
	cachetime = r.Timeunix
	stream, err := timer.BroadcastServerTime(context.Background(), &emptypb.Empty{})
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF || err != nil {
			return err
		}
		accrate = resp.Accrate
		offset = resp.Offset
		cachetime = resp.Timeunix
	}
}

func SetAddress(d string) {
	address = d
}

func Shutdown() {
	conn.Close()
}

func Unix() int64 {
	if cachetime != 0 {
		t := time.Now().Unix() - cachetime
		return time.Now().Unix() + int64(offset) + (t * int64(accrate))
	}
	return time.Now().Unix()
}

func Now() time.Time {
	if cachetime != 0 {
		t := time.Now().Unix() - cachetime
		return time.Unix(time.Now().Unix()+int64(offset)+(t*int64(accrate)), 0)
	}
	return time.Now()
}

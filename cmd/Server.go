package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	tm "github.com/wangyuche/time-manager/proto/timemanager"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var url string
var offset int32
var accrate int32

func init() {
	RootCmd.AddCommand(gettimeCmd)
	RootCmd.AddCommand(settimeCmd)
	gettimeCmd.Flags().StringVarP(&url, "url", "u", "localhost:9000", "Time Manager URL")
	settimeCmd.Flags().StringVarP(&url, "url", "u", "localhost:9000", "Time Manager URL")
	settimeCmd.Flags().Int32VarP(&offset, "offset", "o", 0, "Time offset sec")
	settimeCmd.Flags().Int32VarP(&accrate, "accrate", "a", 0, "Time accrate")
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var RootCmd = &cobra.Command{
	Use:   "time-manager",
	Short: "time-manager",
	Long:  `time-manager controls the time server manager.`,
}

var gettimeCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Server Time",
	Run: func(cmd *cobra.Command, args []string) {
		conn, t := Connect(url)
		defer conn.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		r, err := t.GetServerTime(ctx, &emptypb.Empty{})
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		log.Print(r.Accrate)
		log.Print(r.Offset)
	},
}

var settimeCmd = &cobra.Command{
	Use:   "set",
	Short: "Set Server Time",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set-time")
		conn, t := Connect(url)
		defer conn.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		r, err := t.SetServerTime(ctx, &tm.SetServerTimeReq{Offset: offset, Accrate: accrate})
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		if r.GetStatus() != tm.TimeStatus_success {
			log.Fatalf("Statu Code: %v", r.GetStatus())
		}
	},
}

func Connect(address string) (*grpc.ClientConn, tm.TimeGRPCClient) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn, tm.NewTimeGRPCClient(conn)
}

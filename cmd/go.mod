module github.com/wangyuche/time-manager/cmd

go 1.21.4

require (
	github.com/spf13/cobra v1.8.1
	github.com/wangyuche/time-manager/proto/timemanager v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240823204242-4ba0660f739c // indirect
)

replace github.com/wangyuche/time-manager/proto/timemanager => ../proto/timemanager

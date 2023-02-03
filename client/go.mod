module github.com/wangyuche/time-manager/client

go 1.18

require (
	github.com/wangyuche/time-manager/proto/timemanager v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.52.3
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace github.com/wangyuche/time-manager/proto/timemanager => ../proto/timemanager
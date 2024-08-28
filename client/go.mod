module github.com/wangyuche/time-manager/client

go 1.21.4

require (
	github.com/wangyuche/time-manager/proto/timemanager v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
)

require (
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
)

replace github.com/wangyuche/time-manager/proto/timemanager => ../proto/timemanager

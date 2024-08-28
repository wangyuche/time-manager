

build_mac:
	protoc **/*.proto  --go_out=. \
	--go_opt=paths=source_relative --proto_path=. \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative 

#export PATH="$PATH:$(go env GOPATH)/bin"
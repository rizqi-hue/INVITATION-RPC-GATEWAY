
start:
	go run main.go serve --config file://config.json
	
protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/app/auth/pb/*.proto --go-grpc_opt=require_unimplemented_servers=false

proto:
	protoc api/**/pb/*.proto --go-grpc_out=.
generate_grpc_code:
	protoc \
	--go_out=getProducer \
	--go_opt=paths=source_relative \
	--go-grpc_out=getProducer \
	--go-grpc_opt=paths=source_relative \
	getProducer.proto

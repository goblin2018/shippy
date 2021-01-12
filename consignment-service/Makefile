gen:
	protoc --proto_path=. --go_out=. --micro_out=. \
		proto/consignment/consignment.proto
run:
	docker run -p 50051:50051 \
		-e MICRO_SERVICE_ADDRESS= :50051 \
		shippy-service-consignment
build:
	docker build -t shippy-service-consignment .		
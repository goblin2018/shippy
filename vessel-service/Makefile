gen:
	protoc --proto_path=. --go_out=. --micro_out=. \
	proto/vessel/vessel.proto
build:
	docker build -t shippy-service-vessel .

run:
	docker run shippy-service-vessel
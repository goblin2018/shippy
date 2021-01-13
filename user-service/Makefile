gen:
	protoc --proto_path=. --go_out=. --micro_out=. \
	proto/user/user.proto

build:
	docker build -t shippy-service-user .

run:
	docker run shippy-service-user
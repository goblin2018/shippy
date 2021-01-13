module github.com/goblin2018/shippy/consignment-service

go 1.14

require (
	github.com/goblin2018/shippy/vessel-service v0.0.0-20210112154308-421a8620ea46
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro/v2 v2.9.1
	go.mongodb.org/mongo-driver v1.4.4
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

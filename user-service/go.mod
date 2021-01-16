module github.com/goblin2018/shippy/user-service

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.9.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

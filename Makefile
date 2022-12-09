build:
	rm -rf build && mkdir build && go build -o build/app_bonus -v cmd/main.go
  
run:
	go run cmd/main.go

sw:
	swag init -g cmd/main.go

test:
	go test

migratecreate:
	migrate create -ext sql -dir ./schemes -seq init 

migrateup:
	migrate -path ./schemes -database 'postgres://arch:123456@localhost:5432/newbonus?sslmode=disable' up 

migratedown:
	migrate -path ./schemes -database 'postgres://arch:123456@localhost:5432/newbonus?sslmode=disable' down



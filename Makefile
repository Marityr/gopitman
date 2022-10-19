build:
	rm -rf build && mkdir build && go build -o build/app_bonus -v cmd/main.go
  
run:
	go run cmd/main.go

swag:
	swag init -g cmd/main.go

test:
	go test

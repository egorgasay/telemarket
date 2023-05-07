BINARY_NAME=main.out
 
all: build test

build:
	go build -o ${BINARY_NAME} cmd/bot/main.go

test:
	go test -v cmd/bot/main.go

run:
	go build -o ${BINARY_NAME} cmd/bot/main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
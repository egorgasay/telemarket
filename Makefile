BINARY_NAME=main.out
 
all: build test

build:
	go build -o ${BINARY_NAME} cmd/bot/main.go

test:
	go test -v .

run:
	go build -o ${BINARY_NAME} cmd/bot/main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
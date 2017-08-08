BINARY = api-server

default:
	go build -o ${BINARY}

test:
	go test

clean:
	-rm -f ${BINARY}


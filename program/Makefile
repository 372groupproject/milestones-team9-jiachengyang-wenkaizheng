GOPATH=$(shell pwd)
LIB= ./src/Authentication/authentication.go \
	./src/Core/core.go \
	./src/Core/coreProxy.go \
	./src/Core/coreConnection.go \
	./src/Encryption/encryption.go \
	./src/FileParser/jsonParser.go \
	./src/FileParser/csvParser.go \
	./src/Logging/logging.go

all : mySSLocal mySSServer

mySSServer : src/Server.main/server.go $(LIB)
	GOPATH=$(GOPATH) go build -o mySSServer Server.main

mySSLocal : src/Local.main/local.go src/Local.main/Local/localServerInfo.go $(LIB)
	GOPATH=$(GOPATH) go build -o mySSLocal Local.main

clean:
	rm -f *.a
	rm -f mySSLocal
	rm -f mySSServer
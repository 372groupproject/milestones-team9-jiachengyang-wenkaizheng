all:helloWorld readJson

helloWorld : helloWorld.go
	go build helloWorld.go
	./helloWorld

readJson: readJson.go
	go build readJson.go
	./readJson

clean:
	rm -f helloWorld
	rm -f readJson
	rm -f *.o
all : p2 helloWorld readJson
p2 : p2_feature.go
	go build p2_feature.go
	./p2_feature

helloWorld : helloWorld.go
	go build helloWorld.go
	./helloWorld

readJson: readJson.go
	go build readJson.go
	./readJson

clean:
	rm -f helloWorld
	rm -f readJson
	rm -f p2_feature
	rm -f *.o

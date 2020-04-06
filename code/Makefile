all : p2 helloWorld readJson if for defer switch label
p2 : p2_feature.go
	go build p2_feature.go
	./p2_feature
if : p2_if.go
	go build p2_if.go
	./p2_if
for : p2_for.go
	go build p2_for.go
	./p2_for
defer : p2_defer.go
	go build p2_defer.go
	./p2_defer
switch : p2_switch.go
	go build p2_switch.go
	./p2_switch
label: p2_label.go
	go build p2_label.go
	./p2_label
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

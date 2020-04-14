all : p2_feature helloWorld readJson p2_if p2_for p2_defer p2_switch p2_label p3_string p3_int p3_float p3_list p3_map p3_array p3_byte p3_rune p3_feature p3_struct p3_interface
p2_feature : p2_feature.go 
	go build p2_feature.go
	./p2_feature
p2_if : p2_if.go
	go build p2_if.go
	./p2_if
p2_for : p2_for.go
	go build p2_for.go
	./p2_for
p2_defer : p2_defer.go
	go build p2_defer.go
	./p2_defer
p2_switch : p2_switch.go
	go build p2_switch.go
	./p2_switch
p2_label: p2_label.go
	go build p2_label.go
	./p2_label
helloWorld : helloWorld.go
	go build helloWorld.go
	./helloWorld
p3_string : p3_string.go
	go build p3_string.go
	./p3_string
p3_int : p3_int.go
	go build p3_int.go
	./p3_int
p3_float: p3_float.go
	go build p3_float.go
	./p3_float
p3_list: p3_list.go
	go build p3_list.go
	./p3_list
p3_map: p3_map.go
	go build p3_map.go
	./p3_map
p3_array: p3_array.go
	go build p3_array.go
	./p3_array
readJson: readJson.go
	go build readJson.go
	./readJson
p3_byte: p3_byte.go
	go build p3_byte.go
	./p3_byte
p3_rune: p3_rune.go
	go build p3_rune.go
	./p3_rune
p3_feature: p3_feature.go
	go build p3_feature.go
	./p3_feature README.txt
p3_struct : p3_struct.go
	go build p3_struct.go
	./p3_struct
p3_interface : p3_interface.go
	go build p3_interface.go
	./p3_interface
clean:
	rm -f helloWorld
	rm -f readJson
	rm -f p2_feature
	rm -f p2_if
	rm -f p2_for
	rm -f p2_defer
	rm -f p2_switch
	rm -f p2_label
	rm -f p3_array
	rm -f p3_float
	rm -f p3_int
	rm -f p3_list
	rm -f p3_map
	rm -f p3_string
	rm -f p3_byte
	rm -f p3_rune
	rm -f p3_feature
	rm -f p3_struct
	rm -f p3_interface
	rm -f *.o

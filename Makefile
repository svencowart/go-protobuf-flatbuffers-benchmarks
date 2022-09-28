build-flatbuffer:
	~/dev/open-source/flatbuffers/flatc --go ./message/message.fbs

build-protobuf:
	protoc --go_out=./message ./message/message.proto

build:
	make build-flatbuffer && make build-protobuf

benchmark:
	go test -v -bench=. ./...
.PHONY: proto
proto:
	protoc \
		-I=. \
		-I=$(GOPATH)/src/github.com/protocolbuffers/protobuf/src \
		--go_out=$(GOPATH)/src \
		dft/dft.proto

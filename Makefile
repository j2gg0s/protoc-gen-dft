.PHONY: proto
proto:
	protoc \
		-I=. \
		-I=$(GOPATH)/src/github.com/protocolbuffers/protobuf/src \
		--go_out=$(GOPATH)/src \
		dft/dft.proto

.PHONY: examples-basic
examples-basic:
	protoc \
		-I=examples \
		-I=$(GOPATH)/src/github.com/protocolbuffers/protobuf/src \
		-I=$(GOPATH)/src/github.com/j2gg0s/protoc-gen-dft \
		--go_out=$(GOPATH)/src \
		--dft_out=$(GOPATH)/src \
		basic/everything.proto

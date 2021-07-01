ROOT_PATH              := $(shell pwd)
PB_GO_PATH             := pkg/delivery/grpc/pb
PROTO_PATH             := api/protobuf-spec
SWAGGER_SPEC_PATH      := api/swagger-spec
SEARCH_PROTOFILES      := $(shell find ./api/protobuf-spec -type f -name "*.proto" -printf "%f ")

# Docker images
PROTOC_IMAGE           := znly/protoc



.PHONY: protobuf
protobuf:
	@docker run --rm -v $(ROOT_PATH):/proto $(PROTOC_IMAGE) \
		--go_out=plugins=grpc:/proto/$(PB_GO_PATH) \
		--grpc-gateway_out=logtostderr=true:/proto/$(PB_GO_PATH) \
		--swagger_out=logtostderr=true:/proto/$(SWAGGER_SPEC_PATH) \
		-I=/proto/$(PROTO_PATH) $(SEARCH_PROTOFILES)

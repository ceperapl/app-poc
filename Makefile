PROJECT                := app-poc
IMAGE_VERSION          ?= $(USERNAME)-dev-$(GIT_COMMIT)
DOCKERFILE             := docker/Dockerfile
ROOT_PATH              := $(shell pwd)
PB_GO_PATH             := pkg/delivery/grpc/pb
PROTO_PATH             := api/protobuf-spec
SWAGGER_SPEC_PATH      := api/swagger-spec
SEARCH_PROTOFILES      := $(shell find ./api/protobuf-spec -type f -name "*.proto" -printf "%f ")

SWAGGERUI_PORT         ?= 8089
SWAGGER_JSON           ?= /foo/recipes.swagger.json

# Docker images
PROTOC_IMAGE           := znly/protoc


.PHONY: protobuf
protobuf:
	@docker run --rm -v $(ROOT_PATH):/proto $(PROTOC_IMAGE) \
		--go_out=plugins=grpc:/proto/$(PB_GO_PATH) \
		--grpc-gateway_out=logtostderr=true:/proto/$(PB_GO_PATH) \
		--swagger_out=logtostderr=true:/proto/$(SWAGGER_SPEC_PATH) \
		-I=/proto/$(PROTO_PATH) $(SEARCH_PROTOFILES)
		

.PHONY: image
image: test
	@echo "## Building docker image for $(PROJECT)..."
	@docker build -t $(PROJECT):$(IMAGE_VERSION) -f $(DOCKERFILE) .
	@docker tag $(PROJECT):$(IMAGE_VERSION) $(PROJECT):latest
	@docker image prune -f

.PHONY: test
test: check
	@echo "Test..."

.PHONY: check
check:
	@echo "Check $(PROJECT)"

.PHONY: run
run:
	@docker run -ti --rm --name $(PROJECT) -p 8080:8080 -p 9090:9090 -p 8088:8088 \
		$(PROJECT):latest

.PHONY: run-swaggerui
run-swaggerui:
	@echo "## Running swagger ui in docker..."
	@docker run --rm -p $(SWAGGERUI_PORT):8080 -e SWAGGER_JSON=$(SWAGGER_JSON) \
		-v $(ROOT_PATH)/$(SWAGGER_SPEC_PATH):/foo swaggerapi/swagger-ui

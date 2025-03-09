.PHONY: install-tools protogen clean

PROTO_DIR = rawProtos
OUT_DIR = protos

install-tools:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@export PATH=$$PATH:$(go env GOPATH)/bin

protogen: 
	@mkdir -p $(OUT_DIR)
	@find $(PROTO_DIR) -name "*.proto" | while read protofile; do \
		protoc --proto_path=$(PROTO_DIR) \
		       --go_out=paths=source_relative:$(OUT_DIR) \
		       --go-grpc_out=paths=source_relative:$(OUT_DIR) \
		       $$protofile; \
	done

clean:
	@rm -rf $(OUT_DIR)/*

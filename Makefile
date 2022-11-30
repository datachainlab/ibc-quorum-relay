.PHONY: yrly
yrly:
	go build -o yrly ./relayer

.PHONY: proto-gen
proto-gen:
	@echo "Generating Go files from Protobuf IDL"
	docker run \
		--rm \
		-w /workspace \
		-v $(CURDIR):/workspace \
		tendermintdev/sdk-proto-gen:v0.3 \
		sh ./scripts/protocgen.sh

.PHONY: clean
clean:
	-rm yrly

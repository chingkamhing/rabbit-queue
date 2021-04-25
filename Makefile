# output binary file name
output_name = rabbit-queue

# build the source to native OS and platform
.PHONY: build
build:
	go build -ldflags '-extldflags "-static"' -o $(output_name) *.go

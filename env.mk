APP=r2d2
GOPATH?=/go
REPO_NAME=r2d2
OUTPUT_FILE=./bin/$(APP)
DOCKER_WORKING_DIR=$(GOPATH)/src/github.com/marcelocorreia/$(REPO_NAME)
NAMESPACE=marcelocorreia
IMAGE_GO_GLIDE=marcelocorreia/go-glide-builder:latest
TEST_OUTPUT_DIR=tmp

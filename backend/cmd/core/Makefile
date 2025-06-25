SERVICE_NAME := core
PROJECT_ROOT := $(shell git rev-parse --show-toplevel)
BACKEND_ROOT := $(PROJECT_ROOT)/backend
USECASE_IFACES := $(shell find $(BACKEND_ROOT)/$(SERVICE_NAME)/usecase -maxdepth 1 -type f -name "*.go" | grep -v _test.go | grep -v _impl.go))

build:
	go build -o ../../bin/$(SERVICE_NAME)

mockgen:
	@ $(foreach file,$(USECASE_IFACES), echo $(file) && mockgen -source $(file) -destination $(BACKEND_ROOT)/$(SERVICE_NAME)/usecase/mock/$(shell basename $(file)) || exit;)

bobinstall:
	go install github.com/stephenafamo/bob/gen/bobgen-psql@v0.38.0

bobgen: bobgen-psql -c bobgen.yaml

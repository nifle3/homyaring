.SILENT:

run:
	go build -o ./bin/main-test ./cmd/core/main.go
	./bin/main-test

test:
	go test -v ./...

BUILD_TARGETS := linux-amd64 freebsd-amd64 windows-amd64 darwin-arm64

build:
	$(foreach target,$(BUILD_TARGETS),\
		$(eval OS = $(word 1,$(subst -, ,$(target))))\
		$(eval ARCH = $(word 2,$(subst -, ,$(target))))\
		GOOS=$(OS) GOARCH=$(ARCH) go build -o ./bin/main-$(target)$(if $(filter windows,$(OS)),.exe,) ./cmd/core/main.go;)
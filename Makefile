.PHONY: install

install:
	go build -ldflags="-s" -o mirip cmd/mirip/main.go && \
    mv mirip ${GOPATH}/bin

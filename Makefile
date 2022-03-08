.PHONY: install

install:
	go build -ldflags="-w -s" -o mirip cmd/mirip/main.go && \
    mv mirip ${GOPATH}/bin

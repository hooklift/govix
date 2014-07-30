CGO_CFLAGS:=-I$(CURDIR)/vendor/libvix/include
CGO_LDFLAGS:=-L$(CURDIR)/vendor/libvix -lvixAllProducts -ldl -lpthread

export CGO_CFLAGS CGO_LDFLAGS

build:
	go build

test:
	go test

.PHONY: test build

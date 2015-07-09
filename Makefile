## This Source Code Form is subject to the terms of the Mozilla Public
## License, v. 2.0. If a copy of the MPL was not distributed with this
## file, You can obtain one at http://mozilla.org/MPL/2.0/.

CGO_CFLAGS:=-I$(CURDIR)/vendor/libvix/include -Werror
CGO_LDFLAGS:=-L$(CURDIR)/vendor/libvix -lvixAllProducts -ldl -lpthread

DYLD_LIBRARY_PATH:=$(CURDIR)/vendor/libvix
LD_LIBRARY_PATH:=$(CURDIR)/vendor/libvix

export CGO_CFLAGS CGO_LDFLAGS DYLD_LIBRARY_PATH LD_LIBRARY_PATH

build:
	go build

test:
	go test

deps:
	go get -u github.com/hooklift/govmx

.PHONY: test build deps

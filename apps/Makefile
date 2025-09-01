HOME=$(shell pwd)
BLDOUT=$(HOME)/build

EXT=
ifeq (${GOOS},windows)
    EXT=.exe
endif

APPS = daib taskmanager

all: $(APPS)

.PHONY: $(APPS)

daib:
	go build -o $(BLDOUT) go


clean:
	rm -rf $(BLDOUT)
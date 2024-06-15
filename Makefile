.PHONY: build-docker
all: build
FORCE: ;

SHELL  := env LIBRARY_ENV=$(LIBRARY_ENV) $(SHELL)
LIBRARY_ENV ?= dev

build-docker:
	docker build -t dbo:$(VERSION) .

run-docker:
	docker run --restart=always -v $(LOGDIR):/server/storage/logs -d -p $(PORT):3008 dbo:$(VERSION)

struct:
	db2struct --host localhost -d dbo -t $(TABLE) --package myGoPackage --struct $(STRUCT) -p --user root

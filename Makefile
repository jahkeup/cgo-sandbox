go = go
V = -v
X = -x

.PHONY: build

build:
	mkdir -p build
	go build -o ./build/ ./...

check:
	$(go) build $(V) $(X) ./... cmd/cgo

test:
	$(go) test $(V) ./...

run-all: build
	./ci/run-all.sh ./build

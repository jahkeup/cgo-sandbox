go = go
goimports = goimports
golangci_lint = golangci-lint

V = -v
X = -x

GO_PKG_PATH=github.com/jahkeup/cgo-sandbox

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

fmt:: goimports

goimports imports:
	$(goimports) -w -local $(GO_PKG_PATH) .

lint:
	$(golangci_lint) run

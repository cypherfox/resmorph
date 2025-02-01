.PHONY: build test make-pngs lint lint-ci tidy tidy-ci

build: resmorph

./resmorph: test
	go build -o ./resmorph

test:
	go test -cover ./api/... ./cmd/... # for now: ./internal/... ./pkg/... 

lint:
	$(GOBIN)/golangci-lint run ./...

lint-ci:
	$(GOBIN)/golangci-lint run ./... --out-format=colored-line-number --timeout=5m

generate:
	go generate ./...

tidy:
	go mod tidy

tidy-ci:
	tidied -verbose

build-image:
	docker build -f deploy/docker/Dockerfile .
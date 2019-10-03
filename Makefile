SRC=/go/src/github.com/MrDaar/micromachine

.PHONY: lint
lint:
	docker run --rm \
		-v $(CURDIR):$(SRC) \
		-w $(SRC) \
		golangci/golangci-lint:v1.19 golangci-lint run --skip-dirs=vendor

.PHONY: test
test:
	docker run --rm \
		-v $(CURDIR):$(SRC) \
		-w $(SRC) \
		golang:1.13.1-alpine3.10 go test ./... -failfast -count=1

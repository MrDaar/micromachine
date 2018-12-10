.PHONY: lint
lint:
	docker run --rm \
		-v $(CURDIR):/go/src/github.com/MrDaar/micromachine \
		-w /go/src/github.com/MrDaar/micromachine \
		golangci/golangci-lint golangci-lint run --skip-dirs=vendor

.PHONY: test
test:
	go test ./... -count=1

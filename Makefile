GOBIN = $(PWD)/bin

.PHONY: bench clean parser.gen generate run lint security critic test convey build release precommit.rehooks init

clean:
	rm -rf ./tmp coverage.out

parser.gen:
	go generate ./pkg/sdl/parser

generate:
	go generate ./...

run:
	$(GOBIN)/gow -v -c -e=go,mod,qy run .

lint:
	$(GOBIN)/golangci-lint run ./...

security:
	$(GOBIN)/gosec -quiet yaml ./...

critic:
	$(GOBIN)/gocritic check -enableAll -disable commentedOutCode ./...

test: clean lint security critic
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

bench:
	go test -bench=. -run=^$ -benchmem -count=10 ./... |tee ./benchmarks/$$(date +%Y-%m-%d_%H:%M:%S)_benchmark.txt

convey:
	$(GOPATH)/bin/goconvey -excludedDirs vendor,node_modules,bin,dist,build,.vscode,samples -launchBrowser=false

# install: test
# 	CGO_ENABLED=0 go build -ldflags="-s -w" -o $(GOPATH)/bin/cgapp ./cmd/cgapp/main.go

build: test
	$(GOBIN)/goreleaser --snapshot --skip-publish --rm-dist

release: test
	git tag -a v$(VERSION) -m "$(VERSION)"
	$(GOBIN)/goreleaser --snapshot --skip-publish --rm-dist

precommit.rehooks:
	pre-commit autoupdate
	pre-commit install --install-hooks
	pre-commit install --hook-type commit-msg

init:
	@echo "== 👩‍🌾 init golangci-lint =="
	GOBIN=$(GOBIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0

	@echo ""
	@echo "== init pre-commit setup =="
	pip install pre-commit
	pre-commit install
	make precommit.rehooks

	@echo ""
	@echo "== init goreleaser setup =="
	GOBIN=$(GOBIN) go install github.com/goreleaser/goreleaser@latest

	@echo ""
	@echo "== install convey =="
	go get github.com/smartystreets/goconvey

	@echo ""
	@echo "== install gomock =="
	GOBIN=$(GOBIN) go install github.com/golang/mock/mockgen@v1.6.0

	@echo ""
	@echo "== install gosec =="
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s latest

	@echo ""
	@echo "== install go-critic =="
	GOBIN=$(GOBIN) go install -v github.com/go-critic/go-critic/cmd/gocritic@latest

	@echo ""
	@echo "== install gow =="
	GOBIN=$(GOBIN) go install github.com/mitranim/gow@latest

	@echo ""
	@echo "== install ANTLR4 =="
	curl -o ./bin/antlr-4.11.1-complete.jar https://www.antlr.org/download/antlr-4.11.1-complete.jar

	@echo ""
	@echo "== install air =="
	GOBIN=$(GOBIN) go install github.com/cosmtrek/air@latest

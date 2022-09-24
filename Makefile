init:
	@echo "== 👩‍🌾 init =="
	@echo "== 👩‍🌾 golangci-lint =="
# https://golangci-lint.run/
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0

	@echo "== pre-commit setup =="
	pip install pre-commit
	pre-commit install

	@echo "== goreleaser setup =="
	go install github.com/goreleaser/goreleaser@latest

	@echo "== install ginkgo =="
# go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
# go get github.com/onsi/gomega/...

	@echo "== install gomock =="
# go install github.com/golang/mock/mockgen@v1.6.0

build.snapshot:
	@echo "== 👩‍🌾 Build snapshot =="
	goreleaser release --snapshot --rm-dist

precommit.rehooks:
	pre-commit autoupdate
	pre-commit install --install-hooks
	pre-commit install --hook-type commit-msg

test.ci:
	@echo "== 🦸‍️ ci.tester =="
	go test ./...

test.ui:
	@echo "== 🦸‍️ ui.tester =="
	$GOPATH/bin/goconvey

test:
	@echo "== 🦸‍️ tester =="
	go test ./...

lint.ci:
	@echo "== 🙆 linter.ci =="
	golangci-lint run -v ./... --fix

lint:
	@echo "== 🙆 linter =="
	golangci-lint run ./... --fix

RM_TARGET = $$(echo "$@" | cut -c 2-)
GO_FILES = $$(find . -name '*.go')

GO_MODULES_CMD = $$(for i in $$(ls -1 ./cmd); do echo ./cmd/$$i; done)
GO_MODULES_PKG = $$(for i in $$(ls -1 ./pkg); do echo ./pkg/$$i; done)
GO_MODULES_INTERNAL = $$(for i in $$(ls -1 ./internal); do echo ./internal/$$i; done)
GO_MODULES_TEST = $$(for i in $$(ls -1 ./test); do echo ./test/$$i; done)
GO_MODULES = $(GO_MODULES_CMD) $(GO_MODULES_PKG) $(GO_MODULES_INTERNAL)

h: help
help:
	@echo "Available targets:"
	@echo
	@echo "General:"
	@echo "   h,   help         Show this prompt."
	@echo "   c,   clean        Clean up all generated files."
	@echo
	@echo "Deployment:"
	@echo "   b,  build         Build the project and its executables."
	@echo "   bd, docker-build  Build the docker container."
	@echo "   cu, compose-up    Launch the compose configuration for production."
	@echo "   cd, compose-down  Terminate the compose configuration for production."
	@echo
	@echo "Documentation:"
	@echo "   g,   godoc        Launch documenation server."
	@echo "   m,   mkdocs       Launch documenation server."
	@echo
	@echo "Tests:"
	@echo "   t,   test         Run all unit tests of the project."
	@echo "   bm,  benchmarks   Run all benchmarks of the project."
	@echo "   i,   integration  Run all integration tests of the project."
	@echo
	@echo "Code Quality:"
	@echo "   l,   lint         Run the linter on all project files."
	@echo "   f,   format       Format all .go files in the project."
	@echo
	@echo "Development:"
	@echo "   r,   run          Launch the url_shortener executable."
	@echo "   du, dev-up        Launch the compose configuration for development."
	@echo "   dd, dev-down      Terminate the compose configuration for development."
	@echo "   s,   stats        Show repository stats."

.PHONY: \
	h help \
	c clean _clean \
	b build \
	bd docker-build \
	cu compose-up \
	cd compose-down \
	g godoc \
	m mkdocs \
	t test \
	bm benchmarks \
	i integration \
	l lint \
	f format \
	r run \
	du dev-up \
	dd dev-down \
	s stats \
	build-deps

_clean:
	@echo -e '\033[0;33m== Cleanup Script ==\033[0m'
c: clean
clean:
	@# Files
	@export file='url_shortener-linux-386'; [ -f "$$file" ] && rm "$$file" && echo "removed $$file";:
	@export file='url_shortener-linux-amd64'; [ -f "$$file" ] && rm "$$file" && echo "removed $$file";:
	@export file='url_shortener-windows-386.exe'; [ -f "$$file" ] && rm "$$file" && echo "removed $$file";:
	@export file='url_shortener-windows-amd64.exe'; [ -f "$$file" ] && rm "$$file" && echo "removed $$file";:
	@export file='coverage.html'; [ -f "$$file" ] && rm "$$file" && echo "removed $$file";:
	@export file='coverage.out'; [ -f "$$file" ] && rm "$$file" && echo "removed $$file";:

	@# Directories
	@export file='./api/build'; [ -d "$$file" ] && rm -rf "$$file" && echo "removed $$file";:

b: build
build: build-deps
	env GOOS=linux GOARCH=386 go build ./cmd/url_shortener; mv url_shortener url_shortener-linux-386
	env GOOS=linux GOARCH=amd64 go build ./cmd/url_shortener; mv url_shortener url_shortener-linux-amd64
	env GOOS=windows GOARCH=386 go build ./cmd/url_shortener; mv url_shortener.exe url_shortener-windows-386.exe
	env GOOS=windows GOARCH=amd64 go build ./cmd/url_shortener; mv url_shortener.exe url_shortener-windows-amd64.exe

bd: docker-build
docker-build:
	@docker build -t ghcr.io/timkral5/url_shortener .

cu: compose-up
compose-up:
	@docker compose up -d --build

cd: compose-down
compose-down:
	@docker compose down

g: godoc
godoc:
	@echo Access documentation here: http://localhost:6060/pkg/?m=all
	@godoc

m: mkdocs
mkdocs:
	@mkdocs serve -a 0.0.0.0:3005

t: test
test: build-deps
	@#go test -v -bench=. -coverprofile "coverage.out" $(GO_MODULES)
	@go test -v -coverprofile "coverage.out" $(GO_MODULES)
	@go tool cover -html "coverage.out" -o "coverage.html"

bm: benchmarks
benchmarks: build-deps
	@set -a; source ./.env; set +a; go test -v -bench=. '-run=^$$' $(GO_MODULES) $(GO_MODULES_TEST)

i: integration
integration: build-deps
	@set -a; source ./.env; set +a; go test -v $(GO_MODULES_TEST)

l: lint
lint: api/build/
	golangci-lint run ./...

f: format
format:
	@echo 'Formatting...'
	@for i in $(GO_FILES); do\
		echo "   $$i";\
		go fmt "$$i";\
	done
	@echo 'done.'

r: api/build/ run
run:
	@set -a; source ./.env; set +a; go run ./cmd/url_shortener

du: dev-up
dev-up:
	@docker compose -f dev.compose.yaml up -d

dd: dev-down
dev-down:
	@docker compose -f dev.compose.yaml down

s: stats
stats:
	@echo "== STATS =="
	@echo
	@echo "Languages:"
	@echo "   Go: $$(git ls-files | grep .go\$$ | xargs wc -l | grep total | xargs | cut -d ' ' -f 1) lines"
	@echo "   Markdown: $$(git ls-files | grep .md\$$ | xargs wc -l | grep total | xargs | cut -d ' ' -f 1) lines"

build-deps: api/build/

api/build/:
	cd api;./generate.sh

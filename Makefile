RM_TARGET = $$(echo "$@" | cut -c 2-)
GO_FILES = $$(find . -name '*.go')

GO_MODULES_CMD = $$(for i in $$(ls -1 ./cmd); do echo ./cmd/$$i; done)
GO_MODULES_PKG = $$(for i in $$(ls -1 ./pkg); do echo ./pkg/$$i; done)
GO_MODULES_INTERNAL = $$(for i in $$(ls -1 ./internal); do echo ./internal/$$i; done)
GO_MODULES = $(GO_MODULES_CMD) $(GO_MODULES_PKG) $(GO_MODULES_INTERNAL)

h: help
help:
	@echo "Available targets:"
	@echo "   h, help    Show this prompt."
	@echo "   b, build   Build the project and its executables."
	@echo "   t, test    Run all tests of the project."
	@echo "   l, lint    Run the linter on all project files."
	@echo "   r, run     Launch the url_shortener executable."
	@echo "   f, format  Format all .go files in the project."
	@echo "   s, stats   Show repository stats."
	@echo "   c, clean   Clean up all generated files."

.PHONY: h help b build t test l lint r run f format s stats _clean c clean coverage

b: build
build: url_shortener

t: test
test:
	-@go test -v -coverprofile "coverage.out" $(GO_MODULES)
	-@go tool cover -html "coverage.out" -o "coverage.html"

l: lint
lint:
	golangci-lint run

r: run
run:
	go run ./cmd/url_shortener

url_shortener:
	@echo ''
	go build ./cmd/url_shortener

f: format
format:
	@echo 'Formatting...'
	@for i in $(GO_FILES); do\
		echo "   $$i";\
		go fmt "$$i";\
	done
	@echo 'done.'

s: stats
stats:
	@echo "== STATS =="
	@echo
	@echo "Languages:"
	@echo "   Go: $$(git ls-files | grep .go\$$ | xargs wc -l | grep total | xargs | cut -d ' ' -f 1) lines"
	@echo "   Markdown: $$(git ls-files | grep .md\$$ | xargs wc -l | grep total | xargs | cut -d ' ' -f 1) lines"

_url_shortener _coverage.html _coverage.out:
	@printf "\033[0;34mTrying to remove the $(RM_TARGET) executable...\033[0m "
	@[ ! -f "$(RM_TARGET)" ] &&\
		echo -e '\033[0;31malready removed.\033[0m';:
	@[ -f "$(RM_TARGET)" ] &&\
		rm "$(RM_TARGET)" && echo -e '\033[0;32mdone.\033[0m';:

_clean:
	@echo -e '\033[0;33m== Cleanup Script ==\033[0m'
c: clean

clean: _clean _url_shortener _coverage.html _coverage.out


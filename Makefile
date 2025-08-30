help:
	@echo "Available targets:"
	@echo "   b, build  Build the project and its executables."
	@echo "   t, test   Run all tests of the project."
	@echo "   r, run    Launch the url_shortener executable."

b: build
build: url_shortener

t: test
test:
	go test ./...

r: run
run:
	go run ./cmd/url_shortener

url_shortener:
	go build ./cmd/url_shortener

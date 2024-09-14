.PHONY: build
build: 
	go build -o bin/rss-agg

.PHONY: run
run: build
	./bin/rss-agg

.PHONY: test
test: test
	go test ./...

.PHONY: test-all
test-all:
	@go test -v ./... 

.PHONY: test-bench
test-bench:
	go test -bench=. -v ./... 
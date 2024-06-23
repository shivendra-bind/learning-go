.PHONY: test-all
test-all:
	@gotestsum  

.PHONY: test-bench
test-bench:
	go test -bench=. -v ./... 


.PHONY: test-cover
test-cover:
	go test -cover -v ./... 
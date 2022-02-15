.PHONY: run
run:
	$(info #Running...)
	go run  ./cmd/overload

.PHONY: run-testserver
run-testserver:
	$(info #Running...)
	go run  ./testserver
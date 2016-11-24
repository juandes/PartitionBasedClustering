TESTPKG := $(shell go list ./... | grep -v /vendor/)


test:
	go test $(TESTPKG)
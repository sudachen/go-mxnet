build:
	cd mx; go build

run-tests:
	cd tests; go test | tee /tmp/test-results

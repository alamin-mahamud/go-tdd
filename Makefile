BINARY=go-tdd
# TESTS=go test $(go list ./... | grep -v /vendor/) -cover
TESTS=go test ./...
TESTS_WITH_COVERAGE=go test ./... -cover
TESTS_WITH_BENCHMARK=go test ./... -bench=.


build:
	${TESTS}
	# go build -o ${BINARY} $(go list ./... | grep -v /vendor/)

install:
	${TESTS}
	# go build -o ${BINARY} $(go list ./... | grep -v /vendor/)

unittest:
	${TESTS_WITH_COVERAGE}

benchmark:
	${TESTS_WITH_BENCHMARK}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install unittest benchmark

TESTDIRS=`go list ./...`
BASEDIR=$(shell pwd)

test:
	go test --cover -v $(TESTDIRS)

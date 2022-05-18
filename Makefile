.PHONY: build clean fmt run

GOCMD=GO111MODULE=on go
BINARY=bin/gw2w

build:
	@echo "build go binary..."
	${GOCMD} build ${GOARGS} -o ${BINARY} cmd/main.go
	@echo "build go binary done"

clean:
	@echo "clean go binary..."
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	@echo "clean go binary done"

fmt:
	@echo "fmt go files..."
	${GOCMD} fmt ${GOARGS} ./...
	@echo "fmt go files done"

run:
	@echo "run server..."
	${GOCMD} run ${GOARGS} cmd/main.go || true

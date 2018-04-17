# Main Makefile for ripe-atlas
#
# Copyright 2016 © by Ollivier Robert
#

.PATH= cmd/fetiche:cmd/fetiched:.
GOBIN=   ${GOPATH}/bin

SRCS= common.go types.go \
	cmd/fetiche/main.go \
	cmd/fetiched/main.go

USRC=	 cmd/fetiche/config_unix.go
WSRC=	cmd/fetiche/config_windows.go

CBIN=	fetiche
DBIN=	fetiched
CEXE=	${CBIN}.exe
DEXE=	${DBIN}.exe
XTRAS=	contrib/* README.md

OPTS=	-ldflags="-s -w" -v

BIN=	${CBIN} ${DBIN}
EXE=	${CEXE} ${DEXE}

all: ${CBIN} ${DBIN}

windows:  ${CEXE} ${DEXE}

${BIN}: ${SRCS} ${USRC}
	go build ${OPTS} ./cmd/...

${EXE}: ${SRCS} ${WSRC}
	GOOS=windows go build ${OPTS} ./cmd/...

test:
	go test -v ./...

bench:
	go test -bench=. -benchmem ./...

lint:
	gometalinter ./...

install: ${BIN}
	go install -v ./cmd/...

pkg: ${BIN} ${EXE}
	-/bin/mkdir pkg
	tar cvf - ${CBIN} ${DBIN} ${XTRAS} | xz > pkg/${CBIN}.tar.xz
	zip -r pkg/${CBIN}.zip ${EXE} ${XTRAS}

clean:
	go clean -v ./...
	rm -f ${BIN} ${EXE} pkg/*

push:
	git push --all
	git push --tags

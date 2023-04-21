FROM golang:1.20.3-alpine AS builder

ENV GOOS "linux"
ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOBIN ${GOPATH}/bin
ENV WORK_DIR ${GOPATH}/src

WORKDIR ${WORK_DIR}

RUN apk update && go install github.com/pressly/goose/v3/cmd/goose@v3.6.1

# =================================================================

FROM alpine:3.16.2
ENV WORK_DIR /src
ENV GOOSE_PATH /go/bin/goose

WORKDIR ${WORK_DIR}
RUN apk update && apk add --no-cache make
COPY ./ ./
COPY --from=builder ${GOOSE_PATH} /bin

CMD [ "make", "all-up" ]

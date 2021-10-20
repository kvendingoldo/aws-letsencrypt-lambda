FROM golang:1.15 AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GIT_TERMINAL_PROMPT=1

COPY ./main.go ${GOPATH}/src/github.com/kvendingoldo/aws-letsencrypt-lambda/main.go
WORKDIR ${GOPATH}/src/github.com/kvendingoldo/aws-letsencrypt-lambda
RUN go build -ldflags="-s -w" -o lambda .

FROM scratch
COPY --from=builder go/src/github.com/kvendingoldo/aws-letsencrypt-lambda/server /app/
WORKDIR /app
ENTRYPOINT ["/app/lambda"]
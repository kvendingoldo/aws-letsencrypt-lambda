FROM golang:1.17 AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GIT_TERMINAL_PROMPT=1

COPY ./main.go ./go.mod ./go.sum ${GOPATH}/src/github.com/kvendingoldo/aws-letsencrypt-lambda/
COPY ./internal ${GOPATH}/src/github.com/kvendingoldo/aws-letsencrypt-lambda/internal
WORKDIR ${GOPATH}/src/github.com/kvendingoldo/aws-letsencrypt-lambda
RUN go get ./
RUN go build -ldflags="-s -w" -o lambda .

FROM scratch
COPY --from=builder go/src/github.com/kvendingoldo/aws-letsencrypt-lambda/lambda /app/
WORKDIR /app
ENTRYPOINT ["/app/lambda"]

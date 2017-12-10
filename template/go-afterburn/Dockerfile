FROM golang:1.9.2-alpine3.6

RUN apk add --no-cache curl && \
    curl -sfL \
    https://github.com/openfaas-incubator/of-watchdog/releases/download/0.0.4/of-watchdog > /usr/bin/fwatchdog && \
    chmod +x /usr/bin/fwatchdog

WORKDIR /go/src/handler
COPY . .

# Run a gofmt and exclude all vendored code.
RUN test -z "$(gofmt -l $(find . -type f -name '*.go' -not -path "./vendor/*"))" || { echo "Run \"gofmt -s -w\" on your Golang code"; exit 1; }

RUN CGO_ENABLED=0 GOOS=linux \
    go build --ldflags "-s -w" -a -installsuffix cgo -o handler . && \
    go test $(go list ./... | grep -v /vendor/) -cover

FROM alpine:3.6
RUN apk --no-cache add ca-certificates

# Add non root user
RUN addgroup -S app && adduser -S -g app app
RUN mkdir -p /home/app
RUN chown app /home/app

WORKDIR /home/app

COPY --from=0 /go/src/handler/handler    .
COPY --from=0 /usr/bin/fwatchdog         .

USER app

ENV fprocess="./handler"

CMD ["./fwatchdog"]

FROM golang:alpine as builder
RUN mkdir -p /go/src/github.com/la3mmchen/elastic-cluster-diff
ADD . /go/src/github.com/la3mmchen/elastic-cluster-diff/
WORKDIR /go/src/github.com/la3mmchen/elastic-cluster-diff/
RUN go build -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder  /go/src/github.com/la3mmchen/elastic-cluster-diff/main /app/
WORKDIR /app
ENTRYPOINT [ "./main" ]
CMD [" --help "]
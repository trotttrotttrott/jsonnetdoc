FROM golang:1.14.2-alpine3.11 as builder
WORKDIR /root
COPY *.go go.mod go.sum ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o jsonnetdoc ./

FROM alpine:3.11.6
COPY --from=builder /root/jsonnetdoc /jsonnetdoc/
ENV PATH $PATH:/jsonnetdoc

ENTRYPOINT ["jsonnetdoc"]

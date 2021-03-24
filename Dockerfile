FROM golang:1.15-alpine as build
WORKDIR /app
ADD . .
ENV GOPATH /go
ENV CGO_ENABLED=0
RUN go build .

FROM alpine:latest
COPY --from=build /app/loggy ./loggy
WORKDIR /
RUN chown 65534:65534 loggy
USER 65534:65534
ENTRYPOINT [ "./loggy" ]

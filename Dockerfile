FROM golang:1.23 as build

WORKDIR /go/src
COPY . .

RUN CGO_ENABLED=0 make


FROM gcr.io/distroless/base:nonroot
USER 65532

ENV MONGODB_MAX_IDLE_TIME_MS=600000
ENV MONGODB_SOCKET_TIMEOUT_MS=20000
ENV MONGODB_CONNECT_TIMEOUT_MS=3000
ENV MONGODB_TIMEOUT_MS=10000
ENV MONGODB_MIN_POOL_SIZE=1
ENV MONGODB_MAX_POOL_SIZE=10
ENV MONGODB_LOG_LEVEL=warn

EXPOSE 27016

COPY --from=build /etc/ssl/certs/ca-certificates.crt \
     /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/src/bin/mongobetween /mongobetween

ENTRYPOINT ["/mongobetween"]

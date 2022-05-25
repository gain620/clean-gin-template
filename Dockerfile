FROM golang:1.18-alpine3.15 AS builder
LABEL stage=gobuilder

WORKDIR /build

ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go test -cover -v ./...; \
    go build -ldflags="-s -w" -o /app/main ./cmd/app/app.go; \
    apk update; \
    apk add upx; \
    upx -9 /app/main

# Alpine Version
FROM alpine:3.15
LABEL maintainer="gainchang620@gmail.com"

ENV APP_USER gin
ENV APP_GROUP docker
ENV APP_UID 666
ENV APP_GID 998
ENV TZ Asia/Seoul
ENV PORT 8080

WORKDIR /app

# TODO : FIX
# cp: can't stat '/usr/share/zoneinfo/Asia/Seoul': No such file or directory
RUN mkdir -p /app/log; \
    cp /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo ${TZ} > /etc/timezone; \
    addgroup \
    --gid ${APP_GID} \
    ${APP_GROUP}; \
    adduser \
    --disabled-password \
    --ingroup ${APP_GROUP} \
    --no-create-home \
    --uid ${APP_UID} \
    ${APP_USER}; \
    chown -R ${APP_USER}:${APP_GROUP} /app

#VOLUME ["/app/log", "/app/log"]

COPY --from=builder /app/main /app/main
EXPOSE ${PORT}

USER ${APP_USER}

HEALTHCHECK --interval=10s --timeout=3s CMD curl --silent --fail http://127.0.0.1:${PORT}/healthz || exit 1
ENTRYPOINT ["./main"]


## Distroless Version
#FROM distroless/static:latest
#LABEL maintainer="gainchang620@gmail.com"
#
##ENV TZ Asia/Seoul
#ENV PORT 8000
#WORKDIR /app
#
## host mkdir -p /app/log && chown -R 65532:65532 /app/log
#VOLUME ["/app/log", "/app/log"]
#
#USER nonroot:nonroot
#
#COPY --from=builder --chown=nonroot:nonroot /app/main /app/main
#COPY --from=builder --chown=nonroot:nonroot /app/log /app/log
#
#EXPOSE ${PORT}
#
## copy compiled app
#HEALTHCHECK --interval=10s --timeout=5s CMD curl -f http://127.0.0.1:${PORT}/healthz || exit 1
#ENTRYPOINT ["./main"]
FROM golang AS builder
WORKDIR /
COPY . .
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct" \
    GOHOSTARCH="amd64" \
    GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
RUN go build -o app cmd/main.go

FROM busybox
COPY --from=builder /app /
COPY config.yaml /
CMD ./app




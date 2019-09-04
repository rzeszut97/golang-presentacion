FROM golang:1.12.4 as builder

RUN mkdir /build
COPY . /build

WORKDIR /build

RUN go mod download && \
    export GO111MODULE=on && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from alpine #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/America/Argentina/Buenos_Aires /etc/localtime && \
    apk del tzdata && rm -rf /var/cache/apk/*

WORKDIR /app
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /build/main .

EXPOSE 8000
CMD ["./main"]
FROM golang:alpine as builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /build 
COPY go.mod /build/
COPY go.sum /build/
WORKDIR /build
RUN go mod download
ADD . /build/
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /build/main /app/

CMD ["./main"]

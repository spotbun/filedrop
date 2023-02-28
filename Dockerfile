FROM golang:alpine AS builder
WORKDIR /app
ADD . .
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /app/main
RUN apk add upx
RUN upx --best --lzma /app/main

FROM scratch
COPY --from=builder /app/main /main
EXPOSE 80
ENTRYPOINT ["/main"]

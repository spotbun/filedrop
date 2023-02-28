FROM golang:alpine AS builder
WORKDIR /app
ADD . .
RUN go build -o /app/main

FROM scratch
COPY --from=builder /app/main /main
EXPOSE 80
ENTRYPOINT ["/main"]

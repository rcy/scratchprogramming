FROM golang:1.20-alpine as builder
WORKDIR /work
COPY . .
RUN go build -o app .

FROM alpine:latest
WORKDIR /work
COPY --from=builder /work/app .
EXPOSE 8080
CMD ["./app"]

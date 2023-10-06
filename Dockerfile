FROM golang:1.21.1-alpine as builder

WORKDIR /app

COPY ./ ./

RUN go mod tidy
RUN go build -o ./bin/cobajenkins ./main.go

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/bin/cobajenkins ./

EXPOSE 9010

CMD ./cobajenkins
FROM golang:1.16

WORKDIR /app

COPY . .

RUN go build request_info.go

CMD ["/app/request_info"]
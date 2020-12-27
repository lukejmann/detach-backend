FROM golang:1.14


ENV GO111MODULE=on
ENV DOCKER=1

RUN mkdir -p /detach/src/forwarding-server
WORKDIR /detach/src/forwarding-server

COPY go.mod .

RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o forwarding_server main.go

CMD ["./forwarding_server"]
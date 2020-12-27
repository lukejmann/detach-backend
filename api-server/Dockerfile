FROM golang:1.14


ENV GO111MODULE=on
ENV DOCKER=1

RUN mkdir -p /detach/src/api-server
WORKDIR /detach/src/api-server

COPY go.mod .

RUN go mod download

COPY . .

# RUN mkdir -p /detach/static
# COPY static /detach/static/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server main.go

CMD ["./api_server"]
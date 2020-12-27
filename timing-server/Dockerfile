FROM golang:1.14


ENV GO111MODULE=on
ENV DOCKER=1


RUN mkdir -p /detach/src/timing-server
WORKDIR /detach/src/timing-server



COPY go.mod .
RUN go mod download


RUN cd 

ADD . .

# Trust Notif Cert
# RUN chmod +rwx notifications/certs/GeoTrust_Global_CA.pem
RUN cp notifications/certs/GeoTrust_Global_CA.pem /usr/local/share/ca-certificates/geotrust_root.crt
RUN update-ca-certificates

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o timing_server main.go

CMD ["./timing_server"]
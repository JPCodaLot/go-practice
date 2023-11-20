FROM golang:1.21-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /usr/local/bin/app ./...

EXPOSE 80/tcp
EXPOSE 3122/udp

CMD ["app"]

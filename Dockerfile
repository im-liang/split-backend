FROM golang:1.22

ENV PORT=8090

WORKDIR /app

COPY src/ .

RUN go mod download
RUN go mod verify

RUN go build -o app

EXPOSE 8090

CMD ["./app"]
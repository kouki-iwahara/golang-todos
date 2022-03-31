FROM golang:1.18.0-alpine

WORKDIR /src

COPY ./go.mod ./

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "main.go"]
FROM golang:1.18.0-alpine

WORKDIR /src

COPY ./go.mod ./

RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
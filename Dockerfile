FROM golang:1.23.3
WORKDIR /app

COPY go.mod ./

RUN go mod download  && go mod verify

COPY . .

RUN go build -o main -v cmd/main.go

CMD ["/app/main"]
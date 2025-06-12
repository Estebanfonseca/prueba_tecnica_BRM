FROM golang:1.23.10

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g ./main.go
RUN go build -o main ./

EXPOSE 8000

CMD ["./main"]
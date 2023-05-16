FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o myapigorm

CMD ["./myapigorm"]
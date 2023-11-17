FROM golang:1.19

WORKDIR /app

COPY . /app

RUN go build ./cmd/main/main.go

EXPOSE 8000

CMD ["go", "run", "./cmd/main/main.go"]

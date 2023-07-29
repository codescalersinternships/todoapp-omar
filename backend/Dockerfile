FROM golang:1.20.5

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Install SQLite3 dependencies (gcc and sqlite3)
RUN apt-get update && \
    apt-get install -y gcc && \
    apt-get install -y libsqlite3-dev

RUN go build -a -o app ./cmd/server.go

CMD [ "./app", "-d", "./internal/data/todoapp.db" ]

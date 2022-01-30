FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
COPY ./init/ ./docker-entrypoint-initdb.d/

RUN go build -o /bin/event ./cmd/event

EXPOSE 8000

CMD ["/bin/event"]
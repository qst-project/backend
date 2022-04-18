FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend

EXPOSE 9091

CMD [ "/app/backend" ]
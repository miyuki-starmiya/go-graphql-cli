FROM golang:1.21.6-alpine

WORKDIR /app

# Install dependencies
RUN apk update && \
  apk add make

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make build

CMD ["./main"]

FROM golang:1.14-alpine

WORKDIR /service/
COPY go.* ./
COPY cmd/ cmd/
COPY internal/ internal/
COPY vendor/ vendor/

WORKDIR /service/cmd/api/

EXPOSE 3000

CMD go run main.go
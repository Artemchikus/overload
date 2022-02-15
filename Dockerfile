FROM golang:1.17-alpine3.13 as development

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/cespare/reflex@latest

EXPOSE 8080

CMD reflex -g '*.go' go run ./cmd/overload/main.go --start-service

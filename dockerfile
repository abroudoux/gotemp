from golang:1.25-alpine as builder

workdir /app

env CGO_ENABLED=0 GOOS=linux GOARCH=amd64

copy go.mod go.sum ./

run go mod download

copy . .

run go build -trimpath -gcflags "-l -B" -ldflags "-s -w" -o gotemp ./cmd/main.go

from golang:1.25-alpine as dev

workdir /app

run go install github.com/air-verse/air@v1.61.7

copy go.mod go.sum ./

run go mod download

expose 8080

entrypoint [ "air", "-c", ".air.toml" ]

FROM gcr.io/distroless/static-debian12

workdir /app

copy --from=builder /app/gotemp .

expose 8080

entrypoint [ "./gotemp" ]

FROM golang:1.18 as builder
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
COPY . .

RUN go build -o server ./src/main.go

FROM scratch
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
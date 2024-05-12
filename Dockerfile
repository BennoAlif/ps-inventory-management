FROM golang:1.22 as builder
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
COPY . .

RUN go build -o server ./src/main.go

FROM scratch
COPY --from=builder /app/server /server

# Set environment variables
ENV DB_NAME=$DB_NAME \
    DB_PORT=$DB_PORT \
    DB_HOST=$DB_HOST \
    DB_USERNAME=$DB_USERNAME \
    DB_PASSWORD=$DB_PASSWORD \
    DB_PARAMS=$DB_PARAMS \
    JWT_SECRET=$JWT_SECRET \
    BCRYPT_SALT=$BCRYPT_SALT

ENTRYPOINT ["/server"]

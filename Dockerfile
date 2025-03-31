FROM golang:latest AS build

# setting work directory to /app
WORKDIR /app

# Copy the go mod files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .


# # =========== stage 2 ============

FROM alpine:latest

WORKDIR /app

# Copy main binary and config file 
COPY --from=build /app/main .
COPY --from=build /app/config.yaml .

# Expose the application port
EXPOSE 8080

ENTRYPOINT ["./main"]
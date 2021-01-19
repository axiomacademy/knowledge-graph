FROM golang:1.14-alpine3.11 as builder
LABEL maintainer="Sudharshan"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:3.11
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 3000
CMD ["sh", "-c", "'sleep 20 && ./main'"]

FROM golang:latest as backend
WORKDIR /app
COPY ./go.mod go.sum ./
RUN go mod download
COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=backend /app/main .
EXPOSE 8080
CMD ["./main"]
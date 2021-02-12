FROM golang:1.12.0-alpine3.9
WORKDIR /app
RUN apk add git
COPY go.mod .
COPY go.sum .
RUN go mod download
ADD . .
RUN go build -o main .
EXPOSE 9000
CMD ["/app/main"]

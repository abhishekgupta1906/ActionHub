
FROM golang:1.20-alpine

WORKDIR /app



# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download


COPY . .

COPY .env .env


RUN go build -o main .


EXPOSE 3000


CMD ["./main"]
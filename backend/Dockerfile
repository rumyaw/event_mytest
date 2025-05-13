FROM golang:1.24.2-alpine

RUN apk add --no-cache gcc musl-dev curl

WORKDIR /app
COPY go.mod ./
COPY . .
RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=1 go build -o main .
RUN ls -l main && test -f main || exit 1

EXPOSE 8080
CMD ["./main"]

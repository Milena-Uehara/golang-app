FROM golang:1.21.0

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o go-web-store main.go

# Port
EXPOSE 8000

# Run
ENTRYPOINT ["/app/go-web-store"]

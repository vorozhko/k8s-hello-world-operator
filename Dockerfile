FROM golang:1.23 AS builder

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o hello-world-operator ./cmd/main.go

# Use a distroless base image for the final image
FROM gcr.io/distroless/base

COPY --from=builder /workspace/hello-world-operator /usr/local/bin/hello-world-operator

ENTRYPOINT ["/usr/local/bin/hello-world-operator"]
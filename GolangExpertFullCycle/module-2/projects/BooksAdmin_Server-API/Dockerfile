# Build stage
# Use the latest Alpine-based Golang image as the starting point for building
FROM golang:alpine AS builder

# Install necessary build tools and libraries
# gcc: GNU Compiler Collection, required for compiling C code (which SQLite uses)
# musl-dev: Development files for musl C library
# sqlite-dev: SQLite development libraries
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
# These files define the project's dependencies
COPY go.mod go.sum ./

# Download all dependencies defined in go.mod
# This step is separated to leverage Docker's caching mechanism
RUN go mod download

# Copy the rest of the source code into the container
COPY . .

# Build the Go application
# CGO_ENABLED=1: Enable cgo, required for SQLite
# GOOS=linux: Set the target operating system to Linux
# go build: Compile the Go program
# -ldflags="-w -s": Linker flags
#   -w: Omit DWARF symbol table info, reducing binary size
#   -s: Omit symbol table and debug info, further reducing size
# -o server: Name the output binary 'server'
# cmd/booksAdmin/main.go: Path to the main Go file
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o server cmd/booksAdmin/main.go

# Final stage: Create the runtime image
# Use Alpine Linux as the base for the final image
FROM alpine:latest

# Install runtime dependencies
# ca-certificates: Required for HTTPS connections
# sqlite-libs: SQLite runtime libraries
RUN apk --no-cache add ca-certificates sqlite-libs

# Set the working directory for the final image
WORKDIR /root/

# Copy only the built binary from the builder stage
# This results in a smaller final image
COPY --from=builder /app/server .

# Expose port 8080 for the application
# Note: This is purely informational and doesn't actually open the port
EXPOSE 8080

# Define the command to run when the container starts
CMD ["./server"]
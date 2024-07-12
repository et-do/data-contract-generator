# Use the official Go image as the base image
FROM golang:1.22.4

# Install essential tools
RUN apt-get update && apt-get install -y \
    git \
    vim \
    curl \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

# Install Go tools
RUN go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install github.com/ramya-rao-a/go-outline@latest \
    && go install github.com/acroca/go-symbols@latest \
    && go install github.com/stamblerre/gocode@latest \
    && go install github.com/rogpeppe/godef@latest \
    && go install github.com/zmb3/gogetdoc@latest \
    && go install golang.org/x/tools/cmd/gorename@latest \
    && go install golang.org/x/tools/cmd/goimports@latest \
    && go install golang.org/x/lint/golint@latest \
    && go install golang.org/x/tools/gopls@latest

# Set the current working directory inside the container
WORKDIR /workspace

# Copy go.mod and go.sum to the working directory
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if go.mod and go.sum are not changed
RUN go mod download

# Copy the source code into the container
COPY . .


# Expose the port for the debugger
EXPOSE 2345

# Command to run when the container starts
CMD ["tail", "-f", "/dev/null"]

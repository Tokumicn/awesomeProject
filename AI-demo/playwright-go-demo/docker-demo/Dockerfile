# Stage 1: Modules caching
#FROM golang:1.22 as modules
#COPY . /modules
#WORKDIR /modules


# Stage 2: Build
FROM golang:1.22 AS builder
COPY . /workdir
WORKDIR /workdir
RUN GOPROXY=https://goproxy.cn
RUN go mod tidy
# Install playwright cli with right version for later use
RUN PWGO_VER=$(grep -oE "playwright-go v\S+" /workdir/go.mod | sed 's/playwright-go //g') \
    && go install github.com/playwright-community/playwright-go/cmd/playwright@${PWGO_VER}
# Build your app
RUN GOOS=linux GOARCH=amd64 go build -o /bin/myapp

# Stage 3: Final
FROM ubuntu:noble
COPY --from=builder /go/bin/playwright /bin/myapp /
RUN apt-get update && apt-get install -y ca-certificates tzdata \
    # Install dependencies and all browsers (or specify one)
    && /playwright install --with-deps \
    && rm -rf /var/lib/apt/lists/*
CMD ["/myapp"]
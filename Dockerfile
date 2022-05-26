FROM golang:latest AS BUILD_IMAGE

WORKDIR /tmp/app

ADD go.mod .
ADD go.sum .
RUN go mod download

ADD . .

# Install dependencies including internal packages. 
RUN go get -d -v ./...
RUN go install -v ./...

# Build the main.go
# Cgo allows to use inline C code in Go sources, Cgo links your application dynamically to libc, even if you don't use any inline C
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o omni-cmd -ldflags="-s -w" ./cmd/omnicmd/...
# -ldflags to ignore debug information for smaller binary
RUN apk add ca-certificates


FROM scratch

# Copy certificates and timezone.
COPY --from=BUILD_IMAGE /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=BUILD_IMAGE /usr/share/zoneinfo /usr/share/zoneinfo
# Copy binary
COPY --from=BUILD_IMAGE /tmp/app/omni-cmd /omni-cmd

EXPOSE 8008

CMD ["/omni-cmd"]
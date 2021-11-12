############################
# STEP 1 build executable binary
############################
# golang alpine 1.16.3
FROM golang@sha256:734250bcdf5b6578e9596bf1e52824306469c7f55923b2e13a29e7c03ccd411e as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata build-base && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
WORKDIR /build

# use modules
COPY . .

RUN go mod download
RUN go mod verify

COPY . .

RUN go test ./...

# Build the binary
RUN GOOS=linux go build -ldflags="-linkmode external -extldflags -static" -a -o /build/app ./cmd/task

############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

WORKDIR /dir

# Copy files
COPY --from=builder --chown=appuser:appuser /build /dir

# Use an unprivileged user.
USER appuser:appuser

VOLUME /dir/data

# Run the app binary.
ENTRYPOINT ["/dir/app"]

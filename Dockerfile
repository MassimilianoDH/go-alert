FROM golang:1.19.2-alpine as builder

# Force Go to use the cgo based DNS resolver. This is required to ensure DNS
# queries required to connect to linked containers succeed.
ENV GODEBUG netdns=cgo

# Install dependencies and build the binaries.
RUN apk add --no-cache --update alpine-sdk \
    git \
    gcc \
&&  git clone https://github.com/MassimilianoDH/go-alert /go/src/github.com/MassimilianoDH/go-alert \
&&  cd /go/src/github.com/MassimilianoDH/go-alert \
&&  go build go-alert

# Start a new, final image.
FROM alpine as final

# Add utilities for quality of life.
RUN apk --no-cache add \
    bash \
    ca-certificates

# Copy the binaries from the builder image.
COPY --from=builder /go/src/github.com/MassimilianoDH/go-alert/go-alert .
COPY --from=builder /go/src/github.com/MassimilianoDH/go-alert/templates ./templates

EXPOSE 8080

ENTRYPOINT [ "./go-alert" ]

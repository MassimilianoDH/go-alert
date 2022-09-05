FROM golang:alpine as builder

WORKDIR /builder

# Force Go to use the cgo based DNS resolver. This is required to ensure DNS
# queries required to connect to linked containers succeed.
ENV GODEBUG netdns=cgo

# Install dependencies and build the binaries.
RUN apk add --no-cache --update alpine-sdk \
    git \
    gcc \
&&  git clone https://github.com/MassimilianoDH/telebot-alert \
&&  cd telebot-alert \
&&  go build

# Start a new, final image.
FROM alpine as final

# add non-root user
RUN adduser -S goalert

# Add bash and ca-certs, for quality of life and SSL-related reasons.
RUN apk --no-cache add \
    bash \
    ca-certificates

# Copy the binaries from the builder image.
COPY --from=builder builder/telebot-alert/go-alert /bin/

CMD [ "go-alert" ]
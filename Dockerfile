#################
# Base image
#################
FROM alpine:3.12 as nodejsweekly-base

USER root

RUN addgroup -g 10001 nodejsweekly && \
    adduser --disabled-password --system --gecos "" --home "/home/nodejsweekly" --shell "/sbin/nologin" --uid 10001 nodejsweekly && \
    mkdir -p "/home/nodejsweekly" && \
    chown nodejsweekly:0 /home/nodejsweekly && \
    chmod g=u /home/nodejsweekly && \
    chmod g=u /etc/passwd

ENV USER=nodejsweekly
USER 10001
WORKDIR /home/nodejsweekly

#################
# Builder image
#################
FROM golang:1.15-alpine AS nodejsweekly-builder
RUN apk add --update --no-cache alpine-sdk
WORKDIR /app
COPY . .
RUN make build

#################
# Final image
#################
FROM nodejsweekly-base

COPY --from=nodejsweekly-builder /app/bin/nodejsweekly /usr/local/bin

# Command to run the executable
ENTRYPOINT ["nodejsweekly"]

# Idea of using a docker build came from https://medium.com/@pierreprinetti/the-go-1-11-dockerfile-a3218319d191

# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.11
ARG GO_VERSION=1.11

FROM golang:${GO_VERSION}-alpine AS builder

# Create the user and group files that will be used in the running container to
# run the process as an unprivileged user
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Install the CA certificates for the app to be able to make calls to HTTPS endpoints
# Ref: https://github.com/drone/ca-certs
RUN apk add --no-cache ca-certificates git

# Set the current working dir
WORKDIR $GOPATH/src/github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot

# Copy the code from the local folder to workdir.
COPY ./ ./

# Copy any configuration files there may be
COPY ./config/*.cfg /app/config/

# Build the executable to `/app`. Mark the build as statically linked (fully self-contained).
RUN CGO_ENABLED=0 go build -o /app.exe .

# Final stage: the running container is scratch since the app is just 1 self-contained executable
FROM scratch

# Copy the user and group files from the builder stage
COPY --from=builder /user/group /user/passwd /etc/

# Copy the CA certificates for enabling HTTPS
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the compiled executable from the builder stage
COPY --from=builder /app.exe ./app

# Copy any configuration files there may be
COPY --from=builder /app/config/ ./config/

# The port used by the application
EXPOSE 13337

# Perform any further action as an unprivileged user
USER nobody:nobody

# The env variables needed for the app
ENV LISTENING_PORT=13337
ENV CERTIFICATE_PATH=/etc/letsencrypt/live/tbp.land/fullchain.pem

ENTRYPOINT ["./app"]

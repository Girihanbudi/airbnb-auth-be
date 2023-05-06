FROM golang:1.19-alpine AS builder

ENV GOPATH /app
ENV GOBIN /go/bin
ENV env ${APP_ENV}

RUN mkdir -p /app/src/airbnb-auth-be
WORKDIR /app/src/airbnb-auth-be
ADD . /app/src/airbnb-auth-be

RUN apk update && apk --no-cache add git openssh-client gcc g++ mercurial alpine-sdk
# RUN mkdir -p -m 0600 ~/.ssh && ssh-keyscan bitbucket.org >> ~/.ssh/known_hosts
# RUN git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"

# RUN --mount=type=ssh,id=tm make build
RUN rm -rf ./dst 
RUN CGO_ENABLED=1 go build -o ./dst/bin/app ./cmd/app/*.go

FROM alpine:3.16

LABEL Maintainer="ghanbudi@gmail.com"
LABEL Service="Auth"
LABEL Web-App="Backend"

WORKDIR /app/src/airbnb-auth-be

RUN apk update && apk add --no-cache bash tzdata ca-certificates curl && \
    cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    echo "Asia/Jakarta" > /etc/timezone \
    apk del curl tzdata && \
    rm -rf /tmp/* && \
    rm -rf /var/cache/apk/*

# Copy necessary used file
COPY --from=builder /app/src/airbnb-auth-be/dst/bin/app /app/src/airbnb-auth-be/main
COPY --from=builder /app/src/airbnb-auth-be/environment /app/src/airbnb-auth-be/environment
COPY --from=builder /app/src/airbnb-auth-be/internal/pkg/firebase /app/src/airbnb-auth-be/internal/pkg/firebase
COPY --from=builder /app/src/airbnb-auth-be/internal/pkg/locale /app/src/airbnb-auth-be/internal/pkg/locale

# Set gin mode to release
ENV GIN_MODE release

EXPOSE 8080
CMD /app/src/airbnb-auth-be/main

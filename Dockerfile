FROM golang:1.11

RUN mkdir -p /usr/src/email-me
WORKDIR /usr/src/email-me

ENV CROSSPLATFORMS \
        linux/amd64 linux/386 linux/arm \
        darwin/amd64 darwin/386 \
        freebsd/amd64 freebsd/386 freebsd/arm \
        windows/amd64 windows/386

ENV GOARM 5

CMD set -x \
    && for platform in $CROSSPLATFORMS; do \
            GOOS=${platform%/*} \
            GOARCH=${platform##*/} \
                go build -v -o bin/email-me-${platform%/*}-${platform##*/}; \
    done

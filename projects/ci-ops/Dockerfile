FROM golang:1.13-alpine AS artifacts
WORKDIR /

RUN apk --no-cache add bash curl git build-base

COPY scripts/deps.sh /deps.sh
RUN bash deps.sh

FROM alpine

RUN apk --no-cache add make ncurses bash coreutils curl git python3 py3-pip gradle moreutils gettext
RUN pip install awscli

COPY --from=artifacts /bin/* /usr/local/bin/

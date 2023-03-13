FROM alpine AS artifacts
WORKDIR /

ENV SPRUCE_VERSION=1.25.2

# ensure we have spruce for parsing yaml
RUN apk add curl && curl -sSLo /opt/spruce https://github.com/geofffranks/spruce/releases/download/v${SPRUCE_VERSION}/spruce-linux-amd64 \
  && chmod +x /opt/spruce

FROM ubuntu:kinetic

COPY --from=artifacts /opt/* /usr/local/bin/

ENV GOLANG_VERSION=1.19 \
  DEBIAN_FRONTEND=noninteractive
  # ^ prevent apt from prompting for user input

# configure paths
ENV GOLANG_PATH=/usr/lib/go-${GOLANG_VERSION}/bin
# ensure we can use everything we'll install
ENV PATH=${GOLANG_PATH}:${PATH}:/root/go/bin

# place startup commands (only run when entrypoint is bash --login)
COPY profile.sh /root/.profile

# install all in single layer
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        apt-utils git curl jq make software-properties-common \
        golang-${GOLANG_VERSION}-go \
        postgresql redis \
    && \
  rm -rf /var/lib/apt/lists/* && \
  go install honnef.co/go/tools/cmd/staticcheck@latest && \
# smoke test the services
  bash /root/.profile && \
  redis-cli --scan && \
  PGPASSWORD=secret psql -h localhost -U postgres -l

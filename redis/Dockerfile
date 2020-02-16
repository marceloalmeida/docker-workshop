ARG FROM_REPO=dw-base
ARG FROM_TAG=latest

FROM ${FROM_REPO}:${FROM_TAG}

ARG VCS_REF
LABEL vcs-ref=$VCS_REF \
      maintainer="Marcelo Almeida <me@malmeida.dev>"

ENV \
  DEBIAN_FRONTEND="noninteractive" \
  DIR="/var/lib/redis" \
  LOGFILE="" \
  BIND="0.0.0.0" \
  REQUIREPASS=""

RUN \
  set -e \
  && apt-get update \
  && apt-get install -y --no-install-recommends --no-install-suggests \
    gosu \
    redis \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh

ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]
CMD ["/usr/bin/redis-server", "/etc/redis/redis.conf"]

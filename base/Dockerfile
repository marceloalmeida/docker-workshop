ARG FROM_REPO=debian
ARG FROM_TAG=bullseye-slim

FROM ${FROM_REPO}:${FROM_TAG}

ARG VCS_REF
LABEL vcs-ref=$VCS_REF \
      maintainer="Marcelo Almeida <me@malmeida.dev>"

ENV \
  DEBIAN_FRONTEND="noninteractive"

RUN \
  echo 'Acquire::Languages "none";' > /etc/apt/apt.conf.d/99disable-translations \
  && set -e \
  && apt-get update \
  && apt-get install -y --no-install-recommends --no-install-suggests \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

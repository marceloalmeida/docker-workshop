#!/bin/bash

DIR="${DIR-/var/lib/redis}"
LOGFILE="${LOGFILE-}"
BIND="${BIND-0.0.0.0}"
REQUIREPASS=${REQUIREPASS-}

CONFIGFILE="/etc/redis/redis.conf"

if [[ "$@" = "/usr/bin/redis-server ${CONFIGFILE}" ]]; then
    mkdir -p "$DIR"
    chown -R redis:redis "$DIR"

    sed -i "s/.*daemonize yes/daemonize no/" ${CONFIGFILE}
    sed -i "s/.*bind.*/bind ${BIND}/" ${CONFIGFILE}

    if [ ! -z "${LOGFILE}" ]; then
      sed -i "s/.*logfile.*/logfile ${LOGFILE}/" ${CONFIGFILE}
    else
      sed -i "/.*logfile.*/d" ${CONFIGFILE}
    fi

    if [ ! -z "${REQUIREPASS}" ]; then
      sed -i "s/.*requirepass.*/requirepass ${REQUIREPASS}/" ${CONFIGFILE}
    fi

    exec gosu redis "$@"
fi

exec "$@"

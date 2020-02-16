import time

import redis
from flask import Flask
from os import getenv

app = Flask(__name__)
cache = redis.Redis(host=getenv('REDIS_HOST'), port=getenv('REDIS_PORT'), password=getenv('REDIS_PASSWORD'))


def get_hit_count():
    retries = 5
    while True:
        try:
            return cache.incr('hits')
        except redis.exceptions.ConnectionError as exc:
            if retries == 0:
                raise exc
            retries -= 1
            time.sleep(0.5)

def get_goapp_hit_count():
    retries = 5
    while True:
        try:
            return cache.get('counter')
        except redis.exceptions.ConnectionError as exc:
            if retries == 0:
                raise exc
            retries -= 1
            time.sleep(0.5)


@app.route('/')
def hello():
    count = get_hit_count()
    goapp_count = get_goapp_hit_count()
    return 'Hello World!\n I have been seen {} times.\n But Go app hitted {}\n'.format(count, int(goapp_count))

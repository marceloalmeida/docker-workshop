# Example app in python

```sh
docker build --build-arg VCS_REF=$(git rev-parse HEAD) --tag dw-pythonapp .
```

```sh
docker run -e REDIS_HOST=172.17.0.2 REDIS_PORT=6379 pythonapp
```


## Source
 * https://docs.docker.com/compose/gettingstarted/

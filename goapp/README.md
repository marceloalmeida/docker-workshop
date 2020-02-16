# Example app in golang

```sh
docker build --build-arg VCS_REF=$(git rev-parse HEAD) --tag dw-goapp .
```

```sh
docker run -e REDIS_URL=172.17.0.2:6379 goapp
```

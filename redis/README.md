# Base image

```sh
docker build --build-arg VCS_REF=$(git rev-parse HEAD) --tag dw-redis .
```

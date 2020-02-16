ARG FROM_REPO=python
ARG FROM_TAG=3.7-alpine

FROM ${FROM_REPO}:${FROM_TAG}

ARG VCS_REF
LABEL vcs-ref=$VCS_REF \
      maintainer="Marcelo Almeida <me@malmeida.dev>"

WORKDIR /code

RUN apk add --no-cache gcc musl-dev linux-headers
COPY requirements.txt requirements.txt

RUN pip install -r requirements.txt
COPY . .

ENV \
  FLASK_APP="app.py" \
  FLASK_RUN_HOST="0.0.0.0" \
  FLASK_RUN_PORT="8000" \
  REDIS_HOST="localhost" \
  REDIS_PORT="6379" \
  REDIS_PASSWORD=""

EXPOSE "8000/tcp"

CMD ["flask", "run"]

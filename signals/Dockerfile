FROM golang:alpine AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go build -o main .

FROM alpine
ARG VCS_REF
LABEL vcs-ref=$VCS_REF \
      maintainer="Marcelo Almeida <me@malmeida.dev>"
WORKDIR /app/
COPY --from=builder /app/main .
CMD ["/app/main"]

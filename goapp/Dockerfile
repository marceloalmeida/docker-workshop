FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o main .

FROM scratch
ARG VCS_REF
LABEL vcs-ref=$VCS_REF \
      maintainer="Marcelo Almeida <me@malmeida.dev>"
ENV \
  REDIS_URL="localhost" \
  REDIS_PASSWORD=""
WORKDIR /app/
COPY --from=builder /app/main .
CMD ["./main"]  

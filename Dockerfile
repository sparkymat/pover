FROM golang:1.22 AS builder

RUN apt-get update && apt-get install -y \
  make \
  && rm -rf /var/lib/apt/lists/*

COPY . /app/

WORKDIR /app
RUN go generate ./...
RUN make pover

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
  ca-certificates \
  ruby \
  povray \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/pover /bin/pover

CMD ["/bin/pover"]

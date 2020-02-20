FROM golang:1.13.8 as builder

WORKDIR /src

COPY . /src

RUN cd /src && go build cmd/publicvault/*

FROM alpine

WORKDIR /app

COPY --from=builder /src/publicvault /app

ENTRYPOINT ["/app/publicvault"]
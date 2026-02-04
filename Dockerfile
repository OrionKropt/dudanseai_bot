# -------- Stage 1: Build --------
FROM golang:1.25.3-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN apk add --no-cache make

RUN make test

RUN make build

# -------- Stage 2: Production image --------

FROM alpine:latest

COPY --from=builder /app/bin/dudanseai_bot .
COPY --from=builder '/app/cmd/config.json' ./config/config.json

EXPOSE 8080

ENTRYPOINT ["./dudanseai_bot"]


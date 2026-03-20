# -------- Stage 1: Base --------
FROM golang:1.25.3-alpine AS base
WORKDIR /app
RUN apk add --no-cache make git
COPY go.mod go.sum ./
RUN go mod download

# -------- Stage 2: Build --------
FROM base AS builder
COPY . .
RUN make test
RUN make build

# -------- Stage 3: Dev --------
FROM base AS dev
RUN go install github.com/air-verse/air@v1.64.5
COPY . .
CMD ["air"]

# -------- Stage 4: Production image --------
FROM alpine:3.23.3 as prod
COPY --from=builder /app/bin/dudanseai_bot .
ENTRYPOINT ["./dudanseai_bot"]


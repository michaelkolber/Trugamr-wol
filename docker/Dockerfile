FROM golang:1.22.3-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o woa

#------------------#

FROM gcr.io/distroless/static-debian12 AS release

COPY --from=builder /app/woa /usr/local/bin/woa

USER nonroot:nonroot

ENTRYPOINT ["/usr/local/bin/woa"]
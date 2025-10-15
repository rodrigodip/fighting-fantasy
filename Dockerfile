FROM golang:1.24-bullseye AS build-base

WORKDIR /app 

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

FROM build-base AS build-production

COPY . .

RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o ffantasy ./cmd/ffantasy/

FROM gcr.io/distroless/static-debian12:nonroot

ENV GIN_MODE=release

WORKDIR /

COPY --from=build-production /app/ffantasy ffantasy

EXPOSE 8080

CMD ["/ffantasy"]

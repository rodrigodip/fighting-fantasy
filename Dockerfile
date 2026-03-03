FROM golang:1.25-bookworm AS build-base
# FROM golang:1.24-bullseye AS build-base

WORKDIR /app 

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

FROM build-base AS dev

RUN go install github.com/air-verse/air@latest

COPY . .

CMD ["air", "-c", ".air.toml"]

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
# Copy templates and static files to a known location in the container
COPY --from=build-production /app/internal/interface/web/templates templates
COPY --from=build-production /app/internal/interface/web/static static

EXPOSE 8080

CMD ["/ffantasy"]

# Build
FROM golang:1.21.1 AS build-stage

ADD . /src
WORKDIR /src

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /api .

# Run
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /api /api
EXPOSE 4000

USER nonroot:nonroot

ENTRYPOINT ["/api"]

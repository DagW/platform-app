FROM golang:latest AS build

WORKDIR /src

COPY go.mod .
COPY go.sum .

RUN go mod download && go mod verify
COPY . .

ARG revision
ARG version
RUN make build

FROM alpine:latest AS run

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

COPY --from=build /src/bin/app /app

EXPOSE 8080
CMD ["/app"]

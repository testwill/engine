# base Go image version.
FROM golang:1.13.0-stretch AS build

RUN apt-get update && \
      apt-get install -y jq && \
      apt-get clean && \
      rm -rf /var/lib/apt/lists/*

WORKDIR /project

# install dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
ARG version
ENV version=${version}
RUN ./scripts/build-engine.sh

FROM ubuntu:18.04
RUN apt-get update && \
      apt-get install -y --no-install-recommends ca-certificates=20180409 && \
      apt-get clean && \
      rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=build /project/engine .
CMD ["./engine"]

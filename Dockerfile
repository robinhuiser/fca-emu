#
# Build
FROM golang:1.16 AS build

ARG BUILD_VERSION
ARG BUILD_COMMIT
ARG BUILD_DATE
ARG BUILD_BY

WORKDIR /go/src
COPY server ./server
COPY ent ./ent
COPY util ./util
COPY generator ./generator
COPY api ./api
COPY swagger-ui ./swagger-ui

COPY main.go .
COPY go.mod .
COPY go.sum .
COPY banner.txt .

RUN go get -d -v ./...
RUN go build -o fca-emu -a \
   -ldflags "-linkmode external -extldflags '-static' -s -w -X main.version=${BUILD_VERSION} -X main.commit=${BUILD_COMMIT} -X main.date=${BUILD_DATE} -X main.builtBy=${BUILD_BY}" .


#
# Runtime
FROM scratch AS runtime

COPY --from=build /go/src/fca-emu ./
EXPOSE 8080/tcp
ENTRYPOINT ["/fca-emu"]

FROM golang:1.16 AS build

WORKDIR /go/src
COPY server ./server
COPY ent ./ent

COPY main.go .
COPY go.mod .
COPY go.sum .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o finite-mock-server .

FROM scratch AS runtime
ARG port

COPY --from=build /go/src/finite-mock-server ./
EXPOSE ${port}/tcp
ENTRYPOINT ["./finite-mock-server"]

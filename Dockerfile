FROM golang:1.16 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .
COPY go.mod .
COPY go.sum .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o finite .

FROM scratch AS runtime
COPY --from=build /go/src/finite ./
EXPOSE 8080/tcp
ENTRYPOINT ["./finite"]

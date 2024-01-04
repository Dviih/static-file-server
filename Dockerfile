FROM golang:1.21-alpine3.19

WORKDIR /build

COPY . /build

RUN CGO_ENABLED=0 go build -ldflags='-s -w -extldflags="-static"' -o static-file-server .

FROM scratch

COPY --from=0 /build/static-file-server /static-file-server

CMD ["/static-file-server"]

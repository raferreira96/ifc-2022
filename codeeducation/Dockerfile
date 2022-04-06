FROM golang:1.18-alpine3.15 as build

RUN apk add --no-cache --update git

WORKDIR /go/src/api

COPY . .

RUN go get -d -v && \
    go install -v && \
    go build -ldflags '-w -s' -a -installsuffix cgo -o hello . 

## Multistage Building

FROM scratch
COPY --from=build /go/src/api/hello .
CMD ["./hello"]
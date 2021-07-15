FROM golang:1.16.6 as builder

ENV ROOT /go/src/urlmap-api
WORKDIR ${ROOT}
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/main

FROM scratch as main

COPY --from=builder /go/bin/main /main
CMD ["/main"]
FROM golang:1.20 as builder
ENV ROOT /go/src/urlmap-api
WORKDIR ${ROOT}
COPY go.mod go.sum ./
RUN go mod download
COPY service/ ./service/
COPY pb/ ./pb/
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/main

FROM debian:buster-slim AS runner
ENV TZ Asia/Tokyo
COPY --from=builder /go/bin/main /main
USER nobody
CMD ["/main"]

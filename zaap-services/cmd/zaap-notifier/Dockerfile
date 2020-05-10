FROM golang AS builder

WORKDIR /go/src/app
COPY . .

WORKDIR /go/src/app/zaap-services

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o zaap-notifier cmd/zaap-notifier/main.go

FROM scratch

COPY --from=builder /go/src/app/zaap-services/zaap-notifier /zaap-notifier

ENTRYPOINT ["/zaap-notifier"]

FROM golang AS builder

WORKDIR /go/src/app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o zaap-runner cmd/main.go

FROM scratch

COPY --from=builder /go/src/app/zaap-runner /zaap-runner

ENTRYPOINT ["/zaap-runner"]

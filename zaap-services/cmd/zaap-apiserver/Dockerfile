FROM golang AS builder

WORKDIR /go/src/app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o zaap-apiserver cmd/zaap-apiserver/main.go

FROM scratch

COPY --from=builder /go/src/app/zaap-apiserver /zaap-apiserver

ENTRYPOINT ["/zaap-apiserver"]

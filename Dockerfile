# Build stage
FROM golang:1.16.6-alpine3.14 AS build-env

RUN addgroup -S -g 20002 appgroup \
    && adduser -D -u 20001 -s /sbin/nologin -g 'Application User' appuser -G appgroup appgroup

RUN apk --update --no-cache add ca-certificates

COPY . /go/src/github.com/ernst01/go-start-http-api/
WORKDIR /go/src/github.com/ernst01/go-start-http-api/

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o /go-start-http-api cmd/go-start-http-api/go-start-http-api.go

# Test stage
FROM build-env as test-env 

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go test ./... -v -cover 

# Final stage
FROM scratch

COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=build-env /etc/passwd /etc/passwd
USER 20001

COPY --from=build-env /go-start-http-api .
EXPOSE 8080

ENTRYPOINT ["/go-start-http-api","--logtostderr=true"]
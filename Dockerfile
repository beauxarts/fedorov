FROM golang:alpine as build
RUN apk add --no-cache --update git
ADD . /go/src/app
WORKDIR /go/src/app
RUN go get ./...
RUN go build \
    -a -tags timetzdata \
    -o fv \
    -ldflags="-s -w -X 'github.com/beauxarts/fedorov/cli.GitTag=`git describe --tags --abbrev=0`'" \
    main.go

FROM alpine:latest
COPY --from=build /go/src/app/fv /usr/bin/fv
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 1510
#backups
VOLUME /var/lib/fedorov/backups
#metadata
VOLUME /var/lib/fedorov/metadata
#input
VOLUME /var/lib/fedorov/input
#output
VOLUME /var/lib/fedorov/output
#covers
VOLUME /var/lib/fedorov/covers
#downloads
VOLUME /var/lib/fedorov/downloads
#imported
VOLUME /var/lib/fedorov/_imported

ENTRYPOINT ["/usr/bin/fv"]
CMD ["serve","-port", "1510", "-stderr"]

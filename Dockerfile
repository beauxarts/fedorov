FROM golang:alpine as build
RUN apk add --no-cache --update git
ADD . /go/src/app
WORKDIR /go/src/app
RUN go get ./...
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -tags timetzdata -o fv main.go

FROM scratch
COPY --from=build /go/src/app/fv /usr/bin/fv
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 1510
#root dir
VOLUME /var/lib/fedorov
#redux dir
VOLUME /var/lib/fedorov/_redux
#covers
VOLUME /var/lib/fedorov/covers

ENTRYPOINT ["/usr/bin/fv"]
CMD ["serve","-port", "1510", "-stderr"]

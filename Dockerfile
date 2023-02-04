FROM golang:alpine as build
RUN apk add --no-cache --update git
ADD . /go/src/app
WORKDIR /go/src/app
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fv main.go

FROM scratch
COPY --from=build /go/src/app/fv /usr/bin/fv

EXPOSE 1510
#root folder
VOLUME /var/lib/fedorov

ENTRYPOINT ["/usr/bin/fv"]
CMD ["serve","-port", "1510", "-stderr"]

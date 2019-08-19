FROM golang:latest
RUN mkdir /go/src/greenspot
RUN cd /go/src/greenspot
RUN mkdir api
ADD .. /go/src/greenspot/api
RUN go get github.com/olivere/elastic
RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/justinas/alice
RUN cd /go/src
RUN go install -v greenspot/api/...
CMD ["/go/bin/gspot-api"]
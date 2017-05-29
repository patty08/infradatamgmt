FROM golang

RUN go get -d -v github.com/sebastienmusso/infradatamgmt

WORKDIR /go/src/github.com/sebastienmusso/infradatamgmt
RUN go build -o surikator
RUN cd rooter && go test

CMD ["./surikator"]

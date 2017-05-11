FROM golang

WORKDIR /go/src
RUN go get github.com/sebastienmusso/infradatamgmt

WORKDIR /go/src/github.com/sebastienmusso/infradatamgmt
RUN go build -o rooter
RUN cd rooter && go test

CMD ["./rooter"]

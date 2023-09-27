FROM golang:1.19.3

WORKDIR /app

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

COPY . .

CMD ["tail", "-f", "/dev/null"]
FROM golang:1.9-alpine

WORKDIR /go/src/url-shortener
RUN apk --no-cache add curl git && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY . .
RUN dep ensure

RUN go build -o url-shortener
CMD /go/src/url-shortener/url-shortener
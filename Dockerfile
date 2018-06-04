FROM golang:1.9-alpine

WORKDIR /go/src/app
RUN apk --no-cache add curl git && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY . .
RUN dep ensure

RUN go build -o go-shortener
CMD /go/src/app/go-shortener
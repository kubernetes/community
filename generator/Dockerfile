FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download
RUN go-wrapper install

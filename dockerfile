FROM golang:latest

RUN mkdir /go/src/demo
COPY . /go/src/demo
WORKDIR /go/src/demo
RUN go build main.go
EXPOSE 4500
CMD ./main
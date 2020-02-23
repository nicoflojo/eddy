FROM golang:1.13.8-alpine3.10
RUN mkdir /server   
ADD . /server
WORKDIR /server
RUN go build -o main .
CMD ["/server/main"]
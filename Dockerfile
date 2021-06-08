FROM golang
WORKDIR /go/src/go-test-program
ADD . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080:8080
CMD ["go-test-program"]
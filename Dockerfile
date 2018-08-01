# iron/go:dev is the alpine image with the go tools added
FROM golang:1.8

# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
COPY . /go/src/github.com/jaytailor/news-server
RUN go get -d -v ./...
RUN go install -v ./...

# Expose the application on port 8080
EXPOSE 8080

CMD ./bin/main
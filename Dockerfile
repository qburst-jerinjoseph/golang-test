FROM golang:1.10.0


# Move to project root
WORKDIR /go/src/lazy-go

# Install dependencies
RUN go get -u github.com/golang/dep/... github.com/ivpusic/rerun
COPY Gopkg.* ./
RUN dep ensure -vendor-only -v

# Other non-vendored files
COPY main.go ./
COPY internal internal
COPY config config


# Install server application
RUN go install lazy-go

CMD [ "lazy-go", "serve" ]

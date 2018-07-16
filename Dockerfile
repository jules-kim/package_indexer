FROM golang
ADD . $GOPATH/src/github.com/jules-kim/package_indexer
RUN go install github.com/jules-kim/package_indexer
EXPOSE 8080
ENTRYPOINT ["package_indexer"]
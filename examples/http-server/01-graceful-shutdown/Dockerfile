FROM golang as binaryBuilder
WORKDIR /go/src/app
COPY . .
RUN env GOOS=linux GOARCH=386 go build -v -o main main.go

FROM scratch
WORKDIR /opt/helloworld/
COPY --from=binaryBuilder /go/src/app/main .
CMD ["./main"]

# References
# https://hub.docker.com/_/golang/
# https://docs.docker.com/engine/reference/builder/#dockerignore-file

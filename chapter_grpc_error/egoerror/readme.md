protoc -I ./helloworld --go-errors_out=paths=source_relative:./helloworld ./helloworld/*.proto
protoc -I ./helloworld --go_out=paths=source_relative:./helloworld ./helloworld/*.proto
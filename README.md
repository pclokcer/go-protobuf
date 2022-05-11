### GO-PROTO

#### BUILD
```bash
protoc --go_out=./protos ./protos/*.proto
```

#### INSTALL
```bash
go get -d -v ./...
go install -v ./...
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

#### RUN
```bash
go run main.go
```

#### SAMPLE
```bash
https://tutorialedge.net/golang/go-protocol-buffer-tutorial
```
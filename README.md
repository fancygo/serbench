# serbench
## Golang 几个主流序列化工具的测试, json, xml, protobuf, msgpack

- [encoding/json](http://golang.org/pkg/encoding/json/)
- [encoding/xml](http://golang.org/pkg/encoding/xml/)
- [github.com/tinylib/msgp](http://github.com/tinylib/msgp)
- [github.com/golang/protobuf](http://github.com/golang/protobuf)
- [github.com/gogo/protobuf](http://github.com/gogo/protobuf)

### 测试环境 go 1.10

- msgpack
    ```go
    go get github.com/tinylib/msgp
    go generate
    ```

- protobuf
    ```go
    go get github.com/golang/protobuf
    go generate
    ```
- gogoprotobuf
    ```go
    go get github.com/gogo/protobuf/gogoproto
    go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
    go generate
    ```
   
### 测试命令
    ```
    go test -bench=. -benchmem
    ```

### 性能测试结果

```
    goos: linux
    goarch: amd64
    pkg: github.com/fancygo/serbench
    BenchmarkMarshalJson-2           1000000          1090 ns/op
    BenchmarkUnmarshalJson-2          500000          3059 ns/op
    BenchmarkMarshalXml-2             300000          4691 ns/op
    BenchmarkUnmarshalXml-2           500000          2735 ns/op
    BenchmarkMarshalMsgp-2          10000000           141 ns/op
    BenchmarkUnmarshalMsgp-2         5000000           288 ns/op
    BenchmarkMarshalProto-2          3000000           514 ns/op
    BenchmarkUnmarshalProto-2        2000000           893 ns/op
    BenchmarkMarshalGogo-2          10000000           132 ns/op
    BenchmarkUnmarshalGogo-2         3000000           499 ns/op
    PASS
    ok      github.com/FancyGo/serbench 17.043s
```

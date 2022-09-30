# wechat

微信功能开发，不包含支付

### 生成
`protoc --go_out=plugins=grpc:. proto/wechat/*.proto`


### 开发部署
- `sh cmd/deploy.sh {git branch}`

### test

- `go test -count=1 -v ./test/ -test.run Test_OfficialGetAuthURL`
- `go test -count=1 -v ./test/ -test.run Test_OfficialCodeGetUser`
- `go test -count=1 -v ./test/ -test.run Test_OfficialOpenidGetUser`
- `go test -count=1 -v ./test/ -test.run Test_SendTemplateMessage`
- `go test -count=1 -v ./test/ -test.run Test_OfficialJsSdk`

### 目录结构
- cmd 命令
    - [cmd/server] 服务运行命令 `go run cmd/server/main.go -c config/config.toml`
- config 配置文件
    - [config/config.toml.example] 示例配置文件 (文件以`.toml`结尾)
- internal 项目核心文件 
    - [internal/config] 配置文件解析
    - [internal/rpc] 服务核心代码
- proto `protobuf`相关文件
    - 只需要新建修改 `wechat.proto`文件，其他文件通过命令生成
- util 辅助公用文件
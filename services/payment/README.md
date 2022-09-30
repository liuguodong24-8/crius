# payment

支付服务

### 生成

`protoc --go_out=plugins=grpc:. proto/payment/*.proto`

### test

- `go test -count=1 -v ./test/ -test.run Test_WechatUnifiedOrder`

### 开发部署

- `sh cmd/deploy.sh {git branch}`

### 目录结构

- cmd 命令
  - [cmd/server] 服务运行命令 `go run cmd/server/main.go -c config/config.toml`
- config 配置文件
  - [config/config.toml.example] 示例配置文件 (文件以`.toml`结尾)
- internal 项目核心文件
  - [internal/config] 配置文件解析
  - [internal/rpc] 服务核心代码
- pkgs 放置所有对外提供服务的公用包
  - [pkgs/payment](./pkgs/payment) 该服务对外提供的方法
- proto `protobuf`相关文件
  - 只需要新建修改 `xxx.proto`文件，其他文件通过命令生成
- wiki 文档

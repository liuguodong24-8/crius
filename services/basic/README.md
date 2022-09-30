# basic

基础辅助服务

### 生成
`protoc --go_out=plugins=grpc:. basic.proto`


### 目录结构
- cmd 命令
    - [cmd/server] 服务运行命令 `go run cmd/server/main.go -c config/config.toml`
- config 配置文件
    - [config/config.toml.example] 示例配置文件 (文件以`.toml`结尾)
- internal 项目核心文件 
    - [internal/config] 配置解析代码
    - [internal/rpc] 服务核心代码
- pkgs 放置所有对外提供服务的公用包
    - [pkgs/basic](./pkgs/basic) 改服务对外提供的方法
- proto `protobuf`相关文件
    - 只需要新建修改 `basic.proto`文件，其他文件通过命令生成
- util 辅助公用文件
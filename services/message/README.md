# message

消息服务

### 生成
- `protoc --go_out=plugins=grpc:. proto/message/*.proto`
- `protoc --doc_out=./ --doc_opt=markdown,wiki.md proto/message/*.proto`

#### 注：如果发现如下提示，
```
--go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC

See https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code for more information.
```
#### 请使用新格式
`protoc --go_out=. --go-grpc_out=. proto/message/filename.proto`

---


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
    - [pkgs/message](./pkgs/message) 改服务对外提供的方法
- proto `protobuf`相关文件
    - 只需要新建修改 `message.proto`文件，其他文件通过命令生成
- util 辅助公用文件

### test

- `go test -v ./pkgs/message/test/`
- `go test -v ./pkgs/message/test/ -test.run Test_ListOfficialLink`
- `go test -v ./pkgs/message/test/ -test.run Test_ListMiniprogramLink`
- `go test -v ./pkgs/message/test/ -test.run Test_UpdateWechatTemplateStatus`
- `go test -v ./pkgs/message/test/ -test.run Test_ShowWechatTemplate`
- `go test -v ./pkgs/message/test/ -test.run Test_ListWechatTemplate`
- `go test -v ./pkgs/message/test/ -test.run Test_UpdateWechatTemplate`
- `go test -count=1 -v ./pkgs/message/test/ -test.run Test_BranchTemplateSetting`


## 调用

### 发送短信
```
gitlab.omytech.com.cn/micro-service/message/pkgs/message

type SmsSendRequest struct {
	AreaCode string //区号
	Phone    string // 电话
	Message  string // 内容
	System   string // 调用平台
	Sign     string // 签章
}

client, err := message.NewClient(ctx, message.ClientConfig{Adress : "ip:port"})

res, err := client.SmsSend(ctx, message.SmsSendRequest{
        AreaCode: "86",         // 区号
		Phone:   "18203058023", // 电话号码
		Message: "测试调用",     // 发送内容
		System:  "memeber",     // 调用系统
		Sign:    "chunk",       // 短信签章
	})
```

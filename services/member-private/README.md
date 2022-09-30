# member-account

会员私有服务

### 生成
- `protoc --go_out=plugins=grpc:. proto/member-private/*.proto`
- `protoc --doc_out=./ --doc_opt=markdown,wiki.md proto/member-private/*.proto`

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
    - [pkgs/member-private](./pkgs/member-private) 该服务对外提供的方法
- proto `protobuf`相关文件
    - 只需要新建修改 `message.proto`文件，其他文件通过命令生成
- wiki 文档

### 可用方法
- [优惠方案组](./wiki/优惠方案组.md)
    - ListPromotion (ListPromotionRequest)  优惠方案组列表
    - CreatePromotion (CreatePromotionRequest)  创建优惠方案组
    - UpdatePromotion (UpdatePromotionRequest) 修改优惠方案组
    - ShowPromotion (ShowPromotionRequest) 优惠方案组详情
    - UpdatePromotionStatus (UpdateStatusRequest) 修改优惠方案组状态
- [优惠方案](./wiki/优惠方案.md)
    - ListPromotionOption (ListPromotionOptionRequest)
    - CreatePromotionOption (CreatePromotionOptionRequest) 
    - ShowPromotionOption (ShowPromotionOptionRequest)
    - UpdatePromotionOption (UpdatePromotionOptionRequest)
    - UpdatePromotionOptionStatus (UpdateStatusRequest)

# merchant

商户基础服务

### 生成
- `protoc --go_out=plugins=grpc:. ./proto/merchant-basic/*.proto`
- `protoc --doc_out=./ --doc_opt=markdown,wiki.md proto/merchant-basic/*.proto`

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
    - [internal/model] 数据库操作代码
- pkgs 放置所有对外提供服务的公用包
    - [pkgs/merchant-basic](./pkgs/merchant-basic) 改服务对外提供的方法
- proto `protobuf`相关文件
    - 只需要新建修改 `merchant-basic.proto`文件，其他文件通过命令生成
- util 辅助公用文件

### test
- 新增积分类型 `go test -count=1 -v ./test/ -test.run Test_CreatePointCategory`
- 消费类型列表 `go test -count=1 -v ./test/ -test.run Test_ListConsumeCategory`
- 修改积分类型 `go test -count=1 -v ./test/ -test.run Test_UpdatePointCategory`
- 修改积分类型状态 `go test -count=1 -v ./test/ -test.run Test_UpdatePointCategoryStatus`
- 积分类型详情 `go test -count=1 -v ./test/ -test.run Test_ShowPointCategory`
- 设置积分规则说明 `go test -count=1 -v ./test/ -test.run Test_SetPointRuleDescribe`
- 获取积分规则说明 `go test -count=1 -v ./test/ -test.run Test_GetPointRuleDescribe`
- 创建积分规则 `go test -count=1 -v ./test/ -test.run Test_CreatePointRule`
- 修改积分规则 `go test -count=1 -v ./test/ -test.run Test_UpdatePointRule`
- 修改积分规则状态 `go test -count=1 -v ./test/ -test.run Test_UpdatePointRuleStatus`
- 积分规则列表 `go test -count=1 -v ./test/ -test.run Test_ListPointRule`
- 获取积分规则详情|获取门店积分规则 `go test -count=1 -v ./test/ -test.run Test_ShowPointRule`
- 获取积分规则已设置门店合集 `go test -count=1 -v ./test/ -test.run Test_GetPointRuleAllBranch`
- 用户地址 `test\member_address_test.go`
    - 新增用户地址 `go test -count=1 -v ./test/ -test.run Test_CreateMemberAddress`
    - 修改用户地址 `go test -count=1 -v ./test/ -test.run Test_UpdateMemberAddress`
    - 用户地址列表 `go test -count=1 -v ./test/ -test.run Test_GetMemberAddress`
    - 设置用户默认地址 `go test -count=1 -v ./test/ -test.run Test_SetMemberAddressDefault`
    - 获取用户默认地址列表 `go test -count=1 -v ./test/ -test.run Test_GetMemberDefaultAddress`
    - 删除用户地址 `go test -count=1 -v ./test/ -test.run Test_DeleteMemberAddress`

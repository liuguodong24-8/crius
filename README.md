# Crius

Crius,罗马神话中的农业和丰收女神，罗马十二主神之一

## 微服务

采用 `git submodule`管理，检出项目 `git clone --recursive https://gitlab.omytech.com.cn/micro-service/Crius.git`

|命令| 描述|
| ----  | ----  |
| `git submodule init` | 初始化本地配置文件 |
| `git submodule update` | 从该项目中抓取所有数据并检出父项目中列出的合适的提交(指定的提交)。 |
| `git submodule update --init --recursive` | 初始并更新 |
| `git submodule foreach git pull` | 拉取所有子模块 |
| `git submodule update --remote` | 立即更新 |

### 生成
- `protoc --go_out=plugins=grpc:. crius.proto`
- `protoc --php_out=composer --grpc_out=composer --plugin=protoc-gen-grpc={path}grpc_php_plugin services/basic/proto/basic.proto`
- `protoc --php_out=composer services/basic/proto/basic.proto`


### 日志
```
"gitlab.omytech.com.cn/micro-service/Crius/util"


util.Logger.Info("test")
util.Logger.WithSleuthCode(ctx).Info("with code")
util.Logger.WithSleuthCode(ctx).WithFields("params", logger.Fields{
    "test": 1,
}).Info("test 222")
```

### test
- `go test -count=1 -v ./test/crius/ -test.run Test_DiscoverServer`

### 开发部署

- `sh cmd/deploy.sh {git branch}`

### 目录结构
- cmd 命令
    - [cmd/server] 服务运行命令 `go run cmd/server/main.go -c config/config.toml`
- config 配置文件
    - [config/config.toml.example] 示例配置文件 (文件以`.toml`结尾)
- internal 项目核心文件 
- pkgs 放置所有对外提供服务的公用包
    - [pkgs/params](./pkgs/params.go) 自定义 `map[string]interface`结构体
    - [pkgs/metadata](./pkgs/metadata.go) 共用ctx携带调用链,读取调用链
    - [pkgs/crius](./pkgs/crius) 改服务对外提供的方法
    - [pkgs/logger](./pkgs/logger) 日志
    - [pkgs/redis](./pkgs/redis) redis
    - [pkgs/db](./pkgs/db) 数据库
- proto `protobuf`相关文件
    - 只需要新建修改 `crius.proto`文件，其他文件通过命令生成
- [composer](https://gitlab.omytech.com.cn/micro-service/composer.git)  `proto`生成的对应PHP代码
- services 放置所有服务
    - [services/basic](https://gitlab.omytech.com.cn/micro-service/basic) 基础辅助服务
    - [services/integral-grade](https://gitlab.omytech.com.cn/micro-service/integral-grade) 积分等级服务
    - [services/member-account](https://gitlab.omytech.com.cn/micro-service/member-account) 会员账户服务
    - [services/member-private](https://gitlab.omytech.com.cn/micro-service/member-private) 会员私有服务
    - [services/merchant-basic](https://gitlab.omytech.com.cn/micro-service/merchant-basic) 商户基础服务
    - [services/message](https://gitlab.omytech.com.cn/micro-service/message) 消息服务
    - [services/appointment](https://gitlab.omytech.com.cn/micro-service/appointment) 预约服务
    - [services/payment](https://gitlab.omytech.com.cn/micro-service/payment) 支付服务
- util 辅助公用文件

### 返回说明
```
type result struct {
    ErrorCode int32 //0成功  1失败 xxx其他错误定义
    ErrorMessage string // 错误描述
    Data interface{} //响应数据
}
```

### 代码风格

**<font color=red>单向依赖注入，控制反转</font>**

**<font color=red>日志，日志，日志！！！！</font>**

- 变量名  驼峰  userName  
- 方法名  有意义!!! 
    - GetUsers 获取用户列表
    - ShowUser 查看单个用户
    - UpdateUser 修改用户信息
    - DeleteUser 删除用户
    - CreateUser 新增用户
    
**rpc返回**
```
int32 error_code = 1;
string error_message = 2;
string data = 3; // 如需返回信息，全部放在data里面，总数量用total
```
### 相关文档
- [Uber Go 语言编码规范](https://github.com/uber-go/guide)
- [目录规范](https://github.com/golang-standards/project-layout)
- [代码规范检测工具](https://github.com/mgechev/revive)
- [代码格式化工具](https://github.com/golang/tools/tree/master/cmd/goimports)
- [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)
- [编写可维护Go语言代码建议](https://github.com/llitfkitfk/go-best-practice)
- [gRPC Go](https://github.com/grpc/grpc-go)  
- [gRPC中文文档](http://doc.oschina.net/grpc)
- [gRPC英文文档](https://grpc.io/docs/languages/go/)
- [数据库gorm](https://gorm.io/zh_CN/docs/index.html)
- [redis](https://github.com/gomodule/redigo)
- [日志zap](https://github.com/uber-go/zap)

### php调用示例
```json
{
  "name": "micro/basic",
  "description": "omy crius basic server",
  "require": {
    "php": ">=7.4",
    "ext-grpc": "*",
    "google/protobuf": "^3.15",
    "grpc/grpc": "^1.35",
    "omy/crius" : "dev-master"
  },
  "config": {
    "optimize-autoloader": true,
    "preferred-install": "dist",
    "sort-packages": true
  },
  "extra": {
    "laravel": {
      "dont-discover": []
    }
  },
  "repositories": [
    {
      "type": "path",
      "url": "/home/clover/GolandProjects/Microservice/Crius"
    }
  ],
  "minimum-stability": "dev"
}
```

```php
<?php
require_once __DIR__ . '/vendor/autoload.php';

$client = new \Omy\Crius\BasicServer\BasicServiceClient("127.0.0.1:12345", [
    'credentials' => \Grpc\ChannelCredentials::createInsecure()
]);


var_dump($client);

$req = new \Omy\Crius\BasicServer\HelloRequest();
$req->setMessage("hello");
list($res, $status) = $client->Hello($req)->wait();

var_dump($status);
var_dump($res);
var_dump($res->getMessage());
```

### shell
`sh deploy.sh {分支名，默认dev}`

- `protoc --php_out=composer --proto_path=proto/ proto/crius.proto `
- `protoc --php_out=composer --proto_path=services/basic/ services/basic/proto/basic/*.proto`
- `protoc --php_out=composer --proto_path=services/merchant-basic/ services/merchant-basic/proto/merchant-basic/*.proto`
- `protoc --php_out=composer --proto_path=services/appointment/ services/appointment/proto/appointment/*.proto`
- `protoc --php_out=composer --proto_path=services/message/ services/message/proto/message/*.proto`
- `protoc --php_out=composer --proto_path=services/payment/ services/payment/proto/payment/*.proto`
- `protoc --php_out=composer --proto_path=services/member-account/ services/member-account/proto/member-account/*.proto`
- `protoc --php_out=composer --proto_path=services/member-private/ services/member-private/proto/member-private/*.proto`
- `protoc --php_out=composer --proto_path=services/wechat/ services/wechat/proto/wechat/*.proto`
- `protoc --php_out=composer --proto_path=services/point/ services/point/proto/point/*.proto`

### 子服务版本描述
- [services/crius](https://gitlab.omytech.com.cn/micro-service/Crius.git) 服务中心

| 版本 | 功能 | 描述    |
| ---- | ---- | ---- |
| v0.1.0 | 初始版本 | 初始版本 |
| 0.1.1 | 微信预约 | 微信预约 |
| 0.2.1 | 账户标签修改 | 账户便签逻辑调整 0.2.1 |
| 0.3.0 | 积分商城 | 积分商城 0.1.0 | 

- - [services/basic](https://gitlab.omytech.com.cn/micro-service/basic) 基础辅助服务

| 版本 | 功能 | 描述    |
| ---- | ---- | ---- |
| v0.1.0 | 省市区，区号 | 初始版本 |

---

- [services/member-account](https://gitlab.omytech.com.cn/micro-service/member-account) 会员账户服务

| 版本 | 功能 | 描述    |
| ---- | ---- | ---- |
| v0.1.0 | 会员、卡、账户 | 初始版本 |
| v0.2.1 | 会员、卡、账户 | 卡账户标签修改 |

---

- [services/member-private](https://gitlab.omytech.com.cn/micro-service/member-private) 会员私有服务

| 版本 | 功能 | 描述    |
| ---- | ---- | ---- |
| v0.1.0 | 优惠方案 | 初始版本 |
| v0.2.1 | 优惠方案，交班 | 交班列印 |

---

- [services/merchant-basic](https://gitlab.omytech.com.cn/micro-service/merchant-basic) 商户基础服务

| 版本 | 功能 | 描述    |
| ---- | ---- | ---- |
| v0.1.0 | 门店、员工、角色、权限 | 初始版本(官网) |
| v0.2.0 | 门店、员工、角色、权限 | 门店添加`localtion`字段(预约、会员) |
| v0.2.1 |微信预约 | 微信用户 |
| v0.2.2 |会员账户修改引用 | 门店最近一次营业日 |
| v0.2.3 |积分商城0.1.0 | 用户积分 |

---

- [services/message](https://gitlab.omytech.com.cn/micro-service/message) 消息服务

| 版本 | 功能 | 描述    |
| ---- | ---- | ---- |
| v0.1.0 | 短信发送 | 初始版本 |

---

- [services/appointment](https://gitlab.omytech.com.cn/micro-service/appointment) 预约服务

| 版本 | 功能 | 描述    |
| ---- | ---- | ---- |
| v0.1.0 | 预约，模版、后台预约，来电 | 初始版本 |

---

- [services/payment](https://gitlab.omytech.com.cn/micro-service/payment) 支付服务

| 版本 | 功能 | 描述    |
| ---- | ---- | ---- |
| v0.1.0 | 微信支付 | 初始版本 |

- [services/point-mall](https://gitlab.omytech.com.cn/micro-service/point-mall.git) 积分商城

| 版本 | 功能 | 描述    |
| ---- | ---- | ---- |
| v0.1.0 | 积分商城 | 初始版本 |

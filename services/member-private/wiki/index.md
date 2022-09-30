### 会员私有服务使用

```go
import memberPrivate "gitlab.omytech.com.cn/micro-service/member-private/pkgs/member-private"

client, err := memberPrivate.NewClient(context.Background(), memberPrivate.ClientConfig{Address: "xxxx"})
```

- [优惠方案组](./优惠方案组.md)
- [优惠方案](./优惠方案.md)

#### 可用方法

- ListPromotion (ListPromotionRequest)  优惠方案组列表
- CreatePromotion (CreatePromotionRequest)  创建优惠方案组
- UpdatePromotion (UpdatePromotionRequest) 修改优惠方案组
- ShowPromotion (ShowPromotionRequest) 优惠方案组详情
- UpdatePromotionStatus (UpdateStatusRequest) 修改优惠方案组状态
- ListPromotionOption (ListPromotionOptionRequest)  优惠方案列表
- CreatePromotionOption (CreatePromotionOptionRequest) 新增优惠方案
- ShowPromotionOption (ShowPromotionOptionRequest)  优惠方案详情
- UpdatePromotionOption (UpdatePromotionOptionRequest) 修改优惠方案
- UpdatePromotionOptionStatus (UpdateStatusRequest) 修改优惠方案状态
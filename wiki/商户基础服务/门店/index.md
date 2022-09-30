## 门店管理

### 提供方法
- 列表 `GetBranch`
- 详情 `ShowBranch`
- 删除 `DeleteBranch`
- 新增 `CreateBranch`
- 修改 `PatchBranch`


```sql
//门店表
create table merchant.branch(
    id uuid primary key,
    merchant_id uuid comment '商户ID',
    name varchar(20) not null comment '门店名',
    province_id uuid not null comment '省份',
    city_id uuid not null comment '城市',
    district_id uuid not null comment '区域',
    address varchar(128) comment '详细地址',
    contact_phone varchar(20) comment '联系电话',
    extra json comment '冗余信息'
    create_time timestamp,
    update_time timestamp
)
```


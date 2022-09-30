## 服务列表


- [基础公用服务](./基础公用服务)
- [商户基础服务](./商户基础服务)
- [会员账户服务](./会员账户服务)
- [积分等级服务](./积分等级服务)
- [会员私有服务](./会员私有服务)
- [消息服务](./消息服务)

|  服务名    |  服务内容    |
| ---- | ---- |
|  基础公用服务    |  省市区、区号    |
|   商户基础服务   |   门店、员工、角色   |
|   会员账户服务   |  会员、卡、账户    |
|   积分等级服务   |  积分、等级    |
|   会员私有服务   |   权限、优惠方案   |
|   消息服务   |   短信发送   |


### 操作记录

- 服务内部增加表：记录每个数据库增删改的操作记录
```sql
create table snapshot(
    id uuid primary key,
    staff_id uuid commnent '操作员工',
    sleuth_code varchar(20) comment '调用链唯一码',
    table_name varchar(20) not null comment '表名',
    snapshot jsonb comment '操作后数据表信息' 
)
```

- 系统内部增加表：记录每个模块的操作记录
```sql
create table operation_history (
    id uuid primary key,
    system varchar(20) not null comment '系统名：会员系统',
    module varchar(64) not null comment '门店管理',
    page varchar(64) not null comment '门店列表',
    action varchar(64) not null comment '操作按钮：查询 | 新增 | 编辑 | 修改'
)
```
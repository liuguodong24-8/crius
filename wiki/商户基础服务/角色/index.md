## 角色权限

### 提供方法
- 列表 `GetRole`
- 详情 `Showrole`
- 删除 `DeleteRole`
- 新增 `CreateRole`
- 修改 `PatchRole`
- 修改状态 `UpdateRoleStatus`


### 数据库

```sql
// 角色
create table merchant.role(
    id uuid primary key,
    merchant_id uuif comment '商户ID',
    name varchar(20) not null comment '名字',
    status varchar(10) comment '启用状态opened closed',
    create_time timestamp,
    update_time timestamp
)

```
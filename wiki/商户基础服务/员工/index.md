## 员工

### 提供方法
- 列表 `GetStaff`
- 详情 `ShowStaff`
- 删除 `DeleteStaff`
- 新增 `CreateStaff`
- 修改 `PatchStaff`
- 修改状态 `UpdateStaffStatus`
- 重置密码 `ResetPassword`


### 数据库

```sql
//员工表
create table merchant.staff(
    id uuid primary key,
    merchant_id uuif comment '商户ID',
    name varchar(20) not null comment '名字',
    phone varchar(20) not null comment '电话',
    phone_code varchar(10) comment '区号',
    gender smallint not null default 0 comment '1男 2女',
    status varchar(10) comment '启用状态opened closed',
    code varchar(20) not null comment '编号',
    quit_time timestamp comment '离职时间',
    create_time timestamp,
    update_time timestamp
)

// 员工角色
create table merchant.staff_role(
    id uuid primary key,
    merchant_id uuif comment '商户ID',
    staff_id uuid comment '员工ID',
    role_id uuid comment '角色ID',
    create_time timestamp,
    update_time timestamp
)

// 员工角色
create table merchant.staff_branch(
    id uuid primary key,
    merchant_id uuif comment '商户ID',
    staff_id uuid comment '员工ID',
    branch_id uuid comment '门店ID',
    create_time timestamp,
    update_time timestamp
)
```
## 商户基础服务

**[仓库地址](https://gitlab.omytech.com.cn/micro-service/merchant-basic)**

- [门店](./门店)
- [员工](./员工)
- [角色](./角色)

### 数据库

```sql
//商户表
create table merchant.merchant(
    id uuid primary key,
    name varchar(20) not null comment '商户名',
    create_time timestamp,
    update_time timestamp
)
```
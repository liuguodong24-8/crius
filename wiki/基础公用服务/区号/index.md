## 区号

### 提供方法
- 区号列表 `GetPhoneArea`

### 数据库

```sql
// 区号配置
create table phone_area (
    id uuid primary key,
    area varchar(10) not null comment '描述 中国大陆',
    code varchar(10) not null comment '区号 +86',
    create_time timestamp,
    update_time timestamp
);
```

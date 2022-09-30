## 省市区


### 提供方法
- 省列表 `GetProvince`
- 省详情 `ShowProvince`
- 市列表 `GetCity`
- 市详情 `ShowCity`
- 区列表 `GetDistrict`
- 区详情 `ShowDistrict`

### 数据库

```sql
// 省
create table province (
    id uuid primary key,
    name varchar(64) not null comment '名'
    create_time timestamp,
    update_time timestamp
);

// 市
create table city (
    id uuid primary key,
    name varchar(64) not null comment '名'
    province_id uuid comment '省ID',
    create_time timestamp,
    update_time timestamp
);

// 区
create table district (
    id uuid primary key,
    name varchar(64) not null comment '名'
    province_id uuid comment '省ID',
    city_id uuid comment '市ID'
    create_time timestamp,
    update_time timestamp
);
```
